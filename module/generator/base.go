package generator

import (
	"github.com/gofunct/common/files"
	"github.com/gofunct/common/ui"
	"path/filepath"
	"sort"
	"strings"

	"github.com/gofunct/common/errors"
	"github.com/jessevdk/go-assets"
	"github.com/spf13/afero"
)

type baseGenerator interface {
	Generate(dir string, data interface{}, cfg generationConfig) error
	Destroy(dir string, data interface{}) error
}

func newBaseGenerator(tmplFs *assets.FileSystem, fs afero.Fs, ui ui.UI) baseGenerator {
	return &baseGeneratorImpl{
		tmplFs: tmplFs,
		fs:     fs,
		ui:     ui,
	}
}

type baseGeneratorImpl struct {
	tmplFs *assets.FileSystem
	fs     afero.Fs
	ui     ui.UI
}

type generationConfig struct {
	skipTest bool
}

func (g *baseGeneratorImpl) Generate(dir string, data interface{}, genCfg generationConfig) error {
	for _, tmplPath := range g.sortedEntryPaths() {
		if genCfg.skipTest && strings.HasSuffix(tmplPath, "_test.go.tmpl") {
			continue
		}
		entry := g.tmplFs.Files[tmplPath]
		path, err := files.TemplateString(strings.TrimSuffix(tmplPath, ".tmpl")).Compile(data)
		if err != nil {
			return errors.Wrapf(err, "failed to parse path: %s", path)
		}
		absPath := filepath.Join(dir, path)
		dirPath := filepath.Dir(absPath)

		// create directory if not exists
		if err := files.CreateDirIfNotExists(g.fs, dirPath); err != nil {
			return errors.WithStack(err)
		}

		// generate content
		body, err := files.TemplateString(string(entry.Data)).Compile(data)
		if err != nil {
			return errors.Wrapf(err, "failed to generate %s", path)
		}

		// check existed entries
		st := statusCreate
		if ok, err := afero.Exists(g.fs, absPath); err != nil {
			// TODO: handle an error
			st = statusSkipped
		} else if ok {
			existedBody, err := afero.ReadFile(g.fs, absPath)
			if err != nil {
				// TODO: handle an error
				st = statusSkipped
			}
			if string(existedBody) == body {
				st = statusIdentical
			} else {
				st = statusSkipped
				g.ui.Error(path[1:] + " is conflicted.")
				res := g.ui.Ask("Overwite it?")

				if strings.Contains(res, "n") {
					st = statusSkipped
				} else {
					if strings.Contains(res, "n") {
						st = statusCreate
					}
				}
			}
		}

		// create
		if st.ShouldCreate() {
			err = afero.WriteFile(g.fs, absPath, []byte(body), 0644)
			if err != nil {
				return errors.Wrapf(err, "failed to write %s", path)
			}
		}

		st.Fprint(g.ui, path[1:])
	}

	return nil
}

func (g *baseGeneratorImpl) Destroy(dir string, data interface{}) error {
	for _, tmplPath := range g.sortedEntryPaths() {
		path, err := files.TemplateString(strings.TrimSuffix(tmplPath, ".tmpl")).Compile(data)
		if err != nil {
			return errors.Wrapf(err, "failed to parse path: %s", path)
		}
		absPath := filepath.Join(dir, path)

		st := statusSkipped
		if ok, err := afero.Exists(g.fs, absPath); err != nil {
			g.ui.Error("failed to get " + path[1:])
			return errors.WithStack(err)
		} else if ok {
			err = g.fs.Remove(absPath)
			if err != nil {
				g.ui.Error("failed to remove " + path[1:])
				return errors.WithStack(err)
			}
			st = statusDelete
		}

		st.Fprint(g.ui, path[1:])

		dirPath := filepath.Dir(path)
		absDirPath := filepath.Dir(absPath)
		if ok, err := afero.DirExists(g.fs, absDirPath); err == nil && ok {
			if r, err := afero.Glob(g.fs, filepath.Join(absDirPath, "*")); err == nil && len(r) == 0 {
				err = g.fs.Remove(absDirPath)
				if err != nil {
					g.ui.Error("failed to remove " + dirPath[1:])
					return errors.Wrapf(err, "failed to remove %q", dirPath[1:])
				}
				statusDelete.Fprint(g.ui, dirPath[1:])
			}
		}
	}

	return nil
}

func (g *baseGeneratorImpl) sortedEntryPaths() []string {
	rootFiles := make([]string, 0, len(g.tmplFs.Files))
	tmplPaths := make([]string, 0, len(g.tmplFs.Files))
	for path, entry := range g.tmplFs.Files {
		if entry.IsDir() {
			continue
		}
		if strings.Count(entry.Path[1:], "/") == 0 {
			rootFiles = append(rootFiles, path)
		} else {
			tmplPaths = append(tmplPaths, path)
		}
	}
	sort.Strings(rootFiles)
	sort.Strings(tmplPaths)
	return append(rootFiles, tmplPaths...)
}

package usecase

import (
	"context"
	"github.com/gofunct/gogen/module"
	""github.com/gofunct/common/ui""
	"github.com/izumin5210/gex"
	"github.com/pkg/errors"
)

// InitializeProjectUsecase is an interface to create a new grapi project.
type InitializeProjectUsecase interface {
	Perform(rootDir string, cfg InitConfig) error
	GenerateProject(rootDir, pkgName string) error
	InstallDeps(rootDir string, cfg InitConfig) error
}

// NewInitializeProjectUsecase creates a new InitializeProjectUsecase instance.
func NewInitializeProjectUsecase(ui ui.Menu, generator module.ProjectGenerator, gexCfg *gex.Config) InitializeProjectUsecase {
	return &initializeProjectUsecase{
		ui:        ui,
		generator: generator,
		gexCfg:    gexCfg,
	}
}

type initializeProjectUsecase struct {
	ui        ui.Menu
	generator module.ProjectGenerator
	gexCfg    *gex.Config
}

func (u *initializeProjectUsecase) Perform(rootDir string, cfg InitConfig) error {
	u.ui.Section("Initialize project")

	var err error
	err = u.GenerateProject(rootDir, cfg.Package)
	if err != nil {
		return errors.Wrap(err, "failed to initialize project")
	}

	u.ui.Subsection("Install dependencies")
	err = u.InstallDeps(rootDir, cfg)
	if err != nil {
		return errors.Wrap(err, "failed to execute `dep ensure`")
	}

	return nil
}

func (u *initializeProjectUsecase) GenerateProject(rootDir, pkgName string) error {
	return errors.WithStack(u.generator.GenerateProject(rootDir, pkgName))
}

func (u *initializeProjectUsecase) InstallDeps(rootDir string, cfg InitConfig) error {
	u.gexCfg.WorkingDir = rootDir
	repo, err := u.gexCfg.Create()
	if err == nil {
		spec := cfg.BuildSpec()
		err = repo.Add(
			context.TODO(),
			"github.com/gofunct/gogen"+spec,
			"github.com/golang/protobuf/protoc-gen-go",
			"github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway",
			"github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger",
		)
	}
	return errors.WithStack(err)
}

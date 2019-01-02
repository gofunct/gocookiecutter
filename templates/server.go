package templates

import "text/template"

func ServerTemplate() *template.Template {
	return MustCreateTemplate("server.go", server)
}

var server = `package {{.Go.Package }}

import (
	"context"
{{range .Go.Imports}}
	"{{.}}"
{{- end}}

	{{.PbGo.PackageName}} "{{ .PbGo.PackagePath }}"
)

// {{.Go.ServerName}} is a composite interface of {{.PbGo.PackageName }}.{{.Go.ServerName}} and gogenserver.Server.
type {{.Go.ServerName}} interface {
	{{.PbGo.PackageName }}.{{.Go.ServerName}}
	gogenserver.Server
}

// New{{.Go.ServerName}} creates a new {{.Go.ServerName}} instance.
func New{{.Go.ServerName}}() {{.Go.ServerName}} {
	return &{{.Go.StructName}}{}
}

type {{.Go.StructName}} struct {
}
{{$go := .Go -}}
{{$pbGo := .PbGo -}}
{{- range .Methods}}
func (s *{{$go.StructName}}) {{.Method}}(ctx context.Context, req *{{.RequestGo $pbGo.PackageName}}) (*{{.ResponseGo $pbGo.PackageName}}, error) {
	// TODO: Not yet implemented.
	return nil, status.Error(codes.Unimplemented, "TODO: You should implement it!")
}
{{end -}}
`

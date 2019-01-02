package templates

import "text/template"

func ReviewDogTemplate() *template.Template {

	return MustCreateTemplate("reviewdog", `runner:
  golint:
    cmd: golint $(go list ./... | grep -v /vendor/)
    format: golint
  govet:
    cmd: go vet $(go list ./... | grep -v /vendor/)
    format: govet
  errcheck:
    cmd: errcheck -asserts -ignoretests -blank ./...
    errorformat:
      - "%f:%l:%c:%m"
  wraperr:
    cmd: wraperr ./...
    errorformat:
      - "%f:%l:%c:%m"
  megacheck:
    cmd: megacheck ./...
    errorformat:
      - "%f:%l:%c:%m"
  unparam:
    cmd: unparam ./...
    errorformat:
      - "%f:%l:%c: %m"
`)
}

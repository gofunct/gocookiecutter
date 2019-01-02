package template

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Init23b808cac963edf44a497827f2a6eff5ddac970f = ""
var _Init471b25b0481aeb20d1d7630357406c2aa310950f = "package = \"{{.packageName}}\"\n\n[gogen]\nserver_dir = \"./app/server\"\n\n[protoc]\nprotos_dir = \"./api/protos\"\nout_dir = \"./api\"\nimport_dirs = [\n  \"./api/protos\",\n  \"./vendor/github.com/grpc-ecosystem/grpc-gateway\",\n  \"./vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis\",\n]\n\n  [[protoc.plugins]]\n  name = \"go\"\n  args = { plugins = \"grpc\", paths = \"source_relative\" }\n\n  [[protoc.plugins]]\n  name = \"grpc-gateway\"\n  args = { logtostderr = true, paths = \"source_relative\" }\n\n  [[protoc.plugins]]\n  name = \"swagger\"\n  args = { logtostderr = true }\n"
var _Initd135936e91856b6159ac2eedcf89aa9f07773f82 = "package main\n\nimport (\n\t\"os\"\n\n\t\"google.golang.org/grpc/grpclog\"\n\n\t\"{{ .importPath }}/app\"\n)\n\nfunc main() {\n\tos.Exit(run())\n}\n\nfunc run() int {\n\terr := app.Run()\n\tif err != nil {\n\t\tgrpclog.Errorf(\"server was shutdown with errors: %v\", err)\n\t\treturn 1\n\t}\n\treturn 0\n}\n"
var _Init38e76c5db8962fa825cf2bd8b23a2dc985c4513e = "*.so\n/vendor\n/bin\n/tmp\n"
var _Init50bb4ac2099b3758964058926b3c90524e478a2c = ""
var _Init71ed560e812a4261bc8b56d9feaef4800830e0b7 = ""
var _Initbc4053f4dd26ceb67e4646e8c1d2cc75897c4dd0 = "package app\n\nimport (\n\t\"github.com/gofunct/gogen/pkg/gogenserver\"\n)\n\n// Run starts the gogenserver.\nfunc Run() error {\n\ts := gogenserver.New(\n\t\tgogenserver.WithDefaultLogger(),\n\t\tgogenserver.WithServers(\n\t\t// TODO\n\t\t),\n\t)\n\treturn s.Serve()\n}\n"
var _Init8d21956ba8abe388f964e47be0f7e5d170a2fce5 = ""

// Init returns go-assets FileSystem
var Init = assets.NewFileSystem(map[string][]string{"/api": []string{}, "/api/protos": []string{".keep.tmpl"}, "/api/protos/type": []string{".keep.tmpl"}, "/": []string{"Gopkg.toml.tmpl", ".gitignore.tmpl", "gogen.toml.tmpl"}, "/cmd/server": []string{"run.go.tmpl"}, "/cmd": []string{}, "/app": []string{"run.go.tmpl"}, "/app/server": []string{".keep.tmpl"}}, map[string]*assets.File{
	"/Gopkg.toml.tmpl": &assets.File{
		Path:     "/Gopkg.toml.tmpl",
		FileMode: 0x1ed,
		Mtime:    time.Unix(1545633496, 1545633496000000000),
		Data:     []byte(_Init23b808cac963edf44a497827f2a6eff5ddac970f),
	}, "/gogen.toml.tmpl": &assets.File{
		Path:     "/gogen.toml.tmpl",
		FileMode: 0x1ed,
		Mtime:    time.Unix(1546388440, 1546388440501719432),
		Data:     []byte(_Init471b25b0481aeb20d1d7630357406c2aa310950f),
	}, "/api": &assets.File{
		Path:     "/api",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1545633496, 1545633496000000000),
		Data:     nil,
	}, "/api/protos/type": &assets.File{
		Path:     "/api/protos/type",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1545633496, 1545633496000000000),
		Data:     nil,
	}, "/api/protos/type/.keep.tmpl": &assets.File{
		Path:     "/api/protos/type/.keep.tmpl",
		FileMode: 0x1ed,
		Mtime:    time.Unix(1545633496, 1545633496000000000),
		Data:     []byte(_Init50bb4ac2099b3758964058926b3c90524e478a2c),
	}, "/cmd": &assets.File{
		Path:     "/cmd",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1545633496, 1545633496000000000),
		Data:     nil,
	}, "/cmd/server/run.go.tmpl": &assets.File{
		Path:     "/cmd/server/run.go.tmpl",
		FileMode: 0x1ed,
		Mtime:    time.Unix(1545633496, 1545633496000000000),
		Data:     []byte(_Initd135936e91856b6159ac2eedcf89aa9f07773f82),
	}, "/app/server": &assets.File{
		Path:     "/app/server",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1545633496, 1545633496000000000),
		Data:     nil,
	}, "/.gitignore.tmpl": &assets.File{
		Path:     "/.gitignore.tmpl",
		FileMode: 0x1ed,
		Mtime:    time.Unix(1545633496, 1545633496000000000),
		Data:     []byte(_Init38e76c5db8962fa825cf2bd8b23a2dc985c4513e),
	}, "/app/server/.keep.tmpl": &assets.File{
		Path:     "/app/server/.keep.tmpl",
		FileMode: 0x1ed,
		Mtime:    time.Unix(1545633496, 1545633496000000000),
		Data:     []byte(_Init71ed560e812a4261bc8b56d9feaef4800830e0b7),
	}, "/api/protos": &assets.File{
		Path:     "/api/protos",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1545633496, 1545633496000000000),
		Data:     nil,
	}, "/api/protos/.keep.tmpl": &assets.File{
		Path:     "/api/protos/.keep.tmpl",
		FileMode: 0x1ed,
		Mtime:    time.Unix(1545633496, 1545633496000000000),
		Data:     []byte(_Init8d21956ba8abe388f964e47be0f7e5d170a2fce5),
	}, "/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1546388458, 1546388458307759403),
		Data:     nil,
	}, "/cmd/server": &assets.File{
		Path:     "/cmd/server",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1545633496, 1545633496000000000),
		Data:     nil,
	}, "/app": &assets.File{
		Path:     "/app",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1546388584, 1546388584930205471),
		Data:     nil,
	}, "/app/run.go.tmpl": &assets.File{
		Path:     "/app/run.go.tmpl",
		FileMode: 0x1ed,
		Mtime:    time.Unix(1546388584, 1546388584929530369),
		Data:     []byte(_Initbc4053f4dd26ceb67e4646e8c1d2cc75897c4dd0),
	}}, "")

# GoGen

## Base Commands

```commandline
A dev utitility tool for golang based projects

Usage:
  gogen [command]

Available Commands:
  cobra       cobra cli opts
  docker      docker opts
  git         git opts
  help        Help about any command
  protoc      protoc opts
  template    Template opts

Flags:
  -h, --help     help for gogen
  -t, --toggle   Help message for toggle

```

## Template Commands

```commandline
Template opts

Usage:
  gogen template [command]

Available Commands:
  func        list sprig template functions
  gen         Generate files from template directory

Flags:
  -d, --dest-path string   specify the path to the output directory (default ".")
  -h, --help               help for template

Use "gogen template [command] --help" for more information about a command.

```

## Protoc Commands

```commandline
Usage:
  gogen protoc [command]

Available Commands:
  gen         Generate protobug files

Flags:
  -h, --help              help for protoc
  -p, --protodir string   directory containing protobuf files (default "api")

Use "gogen protoc [command] --help" for more information about a command.

```

## Git Commands

```commandline
Usage:
  gogen git [command]

Available Commands:
  clone       clone a git repo
  save        Save all files in current project

Flags:
  -h, --help         help for git
  -m, --msg string   remote url of target repo (default "default msg")
  -u, --url string   remote url of target repo

Use "gogen git [command] --help" for more information about a command.

```

## Docker Commands

```commandline
Usage:
  gogen docker [command]

Available Commands:
  images      list docker images

Flags:
  -e, --endpoint string   docker endpoint (default "unix:///var/run/docker.sock")
  -h, --help              help for docker

Use "gogen docker [command] --help" for more information about a command.

```

## Cobra Commands

```commandline
cobra cli opts

Usage:
  gogen cobra [command]

Available Commands:
  gen         generate a new cobra based application

Flags:
  -a, --app string       the desired cli app name
  -h, --help             help for cobra
  -v, --version string   the version of the cli application (default "0.1.0")

Use "gogen cobra [command] --help" for more information about a command.

```
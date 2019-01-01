package app

import (
	"github.com/gofunct/gogen/grapiserver"
)

// Run starts the grapiserver.
func Run() error {
	s := grapiserver.New(
		grapiserver.WithDefaultLogger(),
		grapiserver.WithServers(
		// TODO
		),
	)
	return s.Serve()
}


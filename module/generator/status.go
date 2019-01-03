package generator

import (
	""github.com/gofunct/common/ui""
	"github.com/izumin5210/grapi/pkg/cli"
)

type status int

const (
	statusCreate status = iota
	statusDelete
	statusExist
	statusIdentical
	statusConflicted
	statusForce
	statusSkipped
)

var (
	creatableStatusSet = map[status]struct{}{
		statusCreate: {},
		statusForce:  {},
	}
)

func (s status) Fprint(ui ui.Menu, msg string) {
	switch s {
	case statusCreate, statusForce, statusDelete:
		ui.ItemSuccess(msg)
	case statusConflicted:
		ui.ItemFailure(msg)
	default:
		ui.ItemSkipped(msg)
	}
}

func (s status) ShouldCreate() bool {
	_, ok := creatableStatusSet[s]
	return ok
}

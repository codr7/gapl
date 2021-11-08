package ops

import (
	"github.com/codr7/gapl"
)

type Stop struct {}

var STOP Stop

func (self Stop) Error() string {
	return "STOP"
}

func (self Stop) Eval(pc gapl.PC, vm *gapl.VM) (gapl.PC, error) {
	return pc, self
}

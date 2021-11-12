package ops

import (
	"github.com/codr7/gapl"
)

type Stop struct {}

var STOP Stop

func (self Stop) Error() string {
	return "STOP"
}

func (self Stop) Eval(pc gapl.Pc, vm *gapl.Vm) (gapl.Pc, error) {
	return pc, self
}

func (self Stop) String() string {
	return "STOP"
}

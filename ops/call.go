package ops

import (
	"github.com/codr7/gapl"
)

type Call struct {
	target *gapl.Func
	flags gapl.CallFlags
}

func NewCall(target *gapl.Func, flags gapl.CallFlags) *Call {
	return &Call{target: target, flags: flags}
}

func (self Call) Eval(pc gapl.PC, vm *gapl.VM) (gapl.PC, error) {
	return self.target.Eval(self.flags, pc+1, vm)
}

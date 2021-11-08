package ops

import (
	"fmt"
	"github.com/codr7/gapl"
)

type Call struct {
	form gapl.Form
	target *gapl.Func
	flags gapl.CallFlags
}

func NewCall(form gapl.Form, target *gapl.Func, flags gapl.CallFlags) *Call {
	return &Call{form: form, target: target, flags: flags}
}

func (self Call) Eval(pc gapl.PC, vm *gapl.VM) (gapl.PC, error) {
	stack := &vm.State().Stack
	
	if self.flags.Check && !self.target.Applicable(stack) {
		return pc, fmt.Errorf("Not applicable: %v %v", self.target, stack)
	}
	
	return self.target.Call(self.flags, pc+1, vm)
}

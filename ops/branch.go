package ops

import (
	"github.com/codr7/gapl"
)

type Branch struct {
	form gapl.Form
	RightPc gapl.Pc
}

func NewBranch(form gapl.Form, rightPc gapl.Pc) *Branch {
	return &Branch{form: form, RightPc: rightPc}
}

func (self Branch) Eval(pc gapl.Pc, vm *gapl.Vm) (gapl.Pc, error) {
	cond := vm.Pop()

	if cond.True() {
		return pc+1, nil
	}
	
	return self.RightPc, nil
}

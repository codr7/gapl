package ops

import (
	"github.com/codr7/gapl"
)

type Load struct {
	form gapl.Form
	reg gapl.Reg
}

func NewLoad(form gapl.Form, reg gapl.Reg) *Load {
	return &Load{form: form, reg: reg}
}

func (self Load) Eval(pc gapl.Pc, vm *gapl.Vm) (gapl.Pc, error) {
	v := vm.Regs()[self.reg]
	vm.Push(v.Type(), v.Data())
	return pc+1, nil
}

package ops

import (
	"github.com/codr7/gapl"
)

type Ret struct {
	form gapl.Form
}

func NewRet(form gapl.Form) *Ret {
	return &Ret{form: form}
}

func (self *Ret) Eval(pc gapl.Pc, vm *gapl.Vm) (gapl.Pc, error) {
	return vm.EndFrame().RestoreState(vm)
}

func (self *Ret) String() string {
	return "RET"
}


package forms

import (
	"github.com/codr7/gapl"
)

type Id struct {
	gapl.BasicForm
	name string
}

func (self *Id) Emit(in []Form, vm *Vm) ([]Form, error) {
	v := vm.Scope().Find(self.name)

	if v == nil {
		return in, fmt.Errorf("Unknown id: %v", self.name)
	}
	
	return in, v.Emit(self, vm)
}

func (self Id) Val(vm *Vm) *Val {
	return vm.Scope().Find(self.name)
}

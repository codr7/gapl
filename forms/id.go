package forms

import (
	"fmt"
	"github.com/codr7/gapl"
)

type Id struct {
	gapl.BasicForm
	name string
}

func NewId(pos gapl.Pos, name string) *Id {
	self := new(Id)
	self.Init(pos)
	self.name = name
	return self
}

func (self Id) Name() string {
	return self.name
}

func (self *Id) Emit(in []gapl.Form, vm *gapl.Vm) ([]gapl.Form, error) {
	v := vm.Scope().Find(self.name)

	if v == nil {
		return in, fmt.Errorf("Unknown id: %v", self.name)
	}
	
	return v.Emit(self, in, vm)
}

func (self Id) Val(vm *gapl.Vm) *gapl.Val {
	return vm.Scope().Find(self.name).Literal()
}

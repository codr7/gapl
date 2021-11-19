package forms

import (
	"github.com/codr7/gapl"
	"github.com/codr7/gapl/ops"
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

func (self *Id) Name() string {
	return self.name
}

func isDrop(name string) bool {
	for _, c := range name {
		if c != 'd' {
			return false
		}	
	}

	return true
}

func (self *Id) Emit(in []gapl.Form, vm *gapl.Vm) ([]gapl.Form, error) {
	if isDrop(self.name) {
		vm.Emit(ops.NewDrop(self, len(self.name)))
		return in, nil
	}

	v := vm.Scope().Find(self.name)

	if v == nil {
		return in, gapl.NewEEmit(self.Pos(), "Unknown id: %v", self.name)
	}
	
	return v.Emit(self, in, vm)
}

func (self *Id) Val(vm *gapl.Vm) *gapl.Val {
	found := vm.Scope().Find(self.name)

	if found == nil {
		return nil
	}
	
	return found.Literal()
}

func (self Id) String() string {
	return self.name
}

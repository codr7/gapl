package forms

import (
	"github.com/codr7/gapl"
)

type Literal struct {
	gapl.BasicForm
	val gapl.Val
}

func (self *Literal) Emit(in []Form, vm *Vm) ([]Form, error) {
	return in, self.val.Emit(self, vm)
}

func (self Literal) Val(vm *Vm) *Val {
	return &self.val
}

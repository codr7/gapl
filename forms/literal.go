package forms

import (
	"github.com/codr7/gapl"
)

type Literal struct {
	gapl.BasicForm
	val gapl.Val
}

func (self *Literal) Emit(in []gapl.Form, vm *gapl.Vm) ([]gapl.Form, error) {
	return self.val.Emit(self, in, vm)
}

func (self Literal) Val(vm *gapl.Vm) *gapl.Val {
	return &self.val
}

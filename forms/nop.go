package forms

import (
	"github.com/codr7/gapl"
)

type Nop struct {
	gapl.BasicForm
}

func (self *Nop) Emit(in []gapl.Form, vm *gapl.Vm) ([]gapl.Form, error) {
	return in, nil
}

func (self Nop) String() string {
	return "_"
}

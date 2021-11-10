package forms

import (
	"github.com/codr7/gapl"
)

type Nop struct {
	gapl.BasicForm
}

func (self Nop) Emit(in []Form, vm *VM) ([]Form, error) {
	return in, nil
}


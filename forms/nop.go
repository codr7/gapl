package forms

import (
	"github.com/codr7/gapl"
)

type Nop struct {
	gapl.BasicForm
}

func (self Nop) Emit(in Forms, vm *VM) (Forms, error) {
	return in, nil
}


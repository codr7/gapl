package types

import (
	"github.com/codr7/gapl"
)

type Macro struct {
	Basic
}

func (self Macro) EmitVal(v gapl.Val, form gapl.Form, in []gapl.Form, vm *gapl.Vm) ([]gapl.Form, error) {
	return v.Data().(*gapl.Macro).Emit(form, in, vm)
}


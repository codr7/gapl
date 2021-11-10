package types

import (
	"github.com/codr7/gapl"
	"github.com/codr7/gapl/ops"
)

type Func struct {
	Basic
}

func (self *Func) EmitVal(v gapl.Val, form gapl.Form, in []gapl.Form, vm *gapl.Vm) ([]gapl.Form, error) {
	vm.Emit(ops.NewCall(form, v.Data().(*gapl.Func), gapl.CallFlags{Check: true}))
	return in, nil
}


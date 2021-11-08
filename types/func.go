package types

import (
	"github.com/codr7/gapl"
	"github.com/codr7/gapl/ops"
)

type Func struct {
	Basic
}

func (self *Func) DumpVal(v gapl.Val) string {
	return v.Data().(*gapl.Func).Dump()
}

func (self *Func) EmitVal(v gapl.Val, form gapl.Form, in gapl.Forms, vm *gapl.VM) (gapl.Forms, error) {
	vm.Emit(ops.NewCall(form, v.Data().(*gapl.Func), gapl.CallFlags{Check: true}))
	return in, nil
}


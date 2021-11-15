package types

import (
	"github.com/codr7/gapl"
	"github.com/codr7/gapl/ops"
)

type Reg struct {
	Basic
}

func (self *Reg) EmitVal(v gapl.Val, form gapl.Form, in []gapl.Form, vm *gapl.Vm) ([]gapl.Form, error) {
	vm.Emit(ops.NewLoad(form, v.Data().(gapl.Reg), nil))
	return in, nil
}

func (self *Reg) LiteralVal(v gapl.Val) *gapl.Val {
	return nil
}



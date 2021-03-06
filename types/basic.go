package types

import (
	"github.com/codr7/gapl"
	"github.com/codr7/gapl/ops"
)

type Basic struct {
	gapl.BasicType
}

func (self *Basic) EmitVal(v gapl.Val, form gapl.Form, in []gapl.Form, vm *gapl.Vm) ([]gapl.Form, error) {
	vm.Emit(ops.NewPush(form, v.Type(), v.Data()))
	return in, nil
}

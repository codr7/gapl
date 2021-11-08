package types

import (
	"fmt"
	"github.com/codr7/gapl"
	"github.com/codr7/gapl/ops"
)

type Reg struct {
	Basic
}

func (self *Reg) DumpVal(v gapl.Val) string {
	return fmt.Sprintf("Reg(%v)", v.Data())
}

func (self *Reg) EmitVal(v gapl.Val, form gapl.Form, in gapl.Forms, vm *gapl.VM) (gapl.Forms, error) {
	vm.Emit(ops.NewLoad(form, v.Data().(gapl.Reg)))
	return in, nil
}

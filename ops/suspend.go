package ops

import (
	"fmt"
	"github.com/codr7/gapl"
)

type Suspend struct {
	form gapl.Form
	EndPc gapl.Pc
}

func NewSuspend(form gapl.Form, endPc gapl.Pc) *Suspend {
	return &Suspend{form: form, EndPc: endPc}
}

func (self *Suspend) Eval(pc gapl.Pc, vm *gapl.Vm) (gapl.Pc, error) {
	c := new(gapl.Cont).Init(pc+1)
	c.Suspend(vm)
	vm.Push(vm.ContType, c)
	return self.EndPc, nil
}

func (self *Suspend) String() string {
	return fmt.Sprintf("SUSPEND %v", self.EndPc)
}


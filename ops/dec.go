package ops

import (
	"fmt"
	"github.com/codr7/gapl"
)

type Dec struct {
	form gapl.Form
	delta int
}

func NewDec(form gapl.Form, delta int) *Dec {	
	return &Dec{form: form, delta: delta}
}

func (self *Dec) Eval(pc gapl.Pc, vm *gapl.Vm) (gapl.Pc, error) {
	vm.Push(vm.IntType, vm.Pop().Data().(int) - self.delta)
	return pc+1, nil
}

func (self *Dec) String() string {
	return fmt.Sprintf("DEC %v", self.delta)
}

package ops

import (
	"fmt"
	"github.com/codr7/gapl"
)

type Inc struct {
	form gapl.Form
	delta int
}

func NewInc(form gapl.Form, delta int) *Inc {	
	return &Inc{form: form, delta: delta}
}

func (self *Inc) Eval(pc gapl.Pc, vm *gapl.Vm) (gapl.Pc, error) {
	vm.Push(vm.IntType, vm.Pop().Data().(int) + self.delta)
	return pc+1, nil
}

func (self *Inc) String() string {
	return fmt.Sprintf("INC %v", self.delta)
}

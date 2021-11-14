package ops

import (
	"fmt"
	"github.com/codr7/gapl"
)

type Dec struct {
	form gapl.Form
	y int
}

func NewDec(form gapl.Form, y int) *Dec {	
	return &Dec{form: form, y: y}
}

func (self *Dec) Eval(pc gapl.Pc, vm *gapl.Vm) (gapl.Pc, error) {
	vm.Push(vm.IntType, vm.Pop().Data().(int) - self.y)
	return pc+1, nil
}

func (self *Dec) String() string {
	return fmt.Sprintf("DEC %v", self.y)
}

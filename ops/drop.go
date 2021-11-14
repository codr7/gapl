package ops

import (
	"fmt"
	"github.com/codr7/gapl"
)

type Drop struct {
	form gapl.Form
	count int
}

func NewDrop(form gapl.Form, count int) *Drop {
	return &Drop{form: form, count: count}
}

func (self *Drop) Eval(pc gapl.Pc, vm *gapl.Vm) (gapl.Pc, error) {
	vm.Stack().Drop(self.count)
	return pc+1, nil
}

func (self *Drop) String() string {
	return fmt.Sprintf("DROP %v", self.count)
}

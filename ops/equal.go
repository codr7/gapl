package ops

import (
	"fmt"
	"github.com/codr7/gapl"
)

type Equal struct {
	form gapl.Form
	Left, Right *gapl.Val
}

func NewEqual(form gapl.Form, left, right *gapl.Val) *Equal {	
	return &Equal{form: form, Left: left, Right: right}
}

func (self *Equal) Eval(pc gapl.Pc, vm *gapl.Vm) (gapl.Pc, error) {
	var left, right gapl.Val
	
	if self.Left == nil {
		left = vm.Pop()
	} else {
		left = *self.Left
	}

	if self.Right == nil {
		right = vm.Pop()
	} else {
		right = *self.Right
	}

	vm.Push(vm.BoolType, left.Equal(right))
	return pc+1, nil
}

func (self *Equal) String() string {
	return fmt.Sprintf("EQUAL %v %v", self.Left, self.Right)
}

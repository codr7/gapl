package ops

import (
	"github.com/codr7/gapl"
)

type Push struct {
	val gapl.Val
}

func NewPush(_type gapl.Type, data interface{}) *Push {
	return &Push{val: gapl.NewVal(_type, data)}
}

func (self Push) Eval(pc gapl.PC, vm *gapl.VM) gapl.PC {
	vm.Push(self.val.Type(), self.val.Data())
	return pc+1
}

package ops

import (
	"github.com/codr7/gapl"
)

type Push struct {
	Val gapl.Val
}

func (self Push) Eval(pc gapl.PC, vm *gapl.VM) gapl.PC {
	vm.Push(self.val)
	return pc+1
}

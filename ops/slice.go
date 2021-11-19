package ops

import (
	"fmt"
	"github.com/codr7/gapl"
)

type Slice struct {
	form gapl.Form
	count int
}

func NewSlice(form gapl.Form, count int) *Slice {
	return &Slice{form: form, count: count}
}

func (self *Slice) Eval(pc gapl.Pc, vm *gapl.Vm) (gapl.Pc, error) {
	out := make(gapl.Slice, self.count)

	for i := 0; i < self.count; i++ {
		out[i] = vm.Pop()
	}

	vm.Push(vm.SliceType, out)
	return pc+1, nil
}

func (self *Slice) String() string {
	return fmt.Sprintf("SLICE %v", self.count)
}

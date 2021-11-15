package ops

import (
	"fmt"
	"github.com/codr7/gapl"
)

type Resume struct {
	form gapl.Form
}

func NewResume(form gapl.Form) *Resume {
	return &Resume{form: form}
}

func (self *Resume) Eval(pc gapl.Pc, vm *gapl.Vm) (gapl.Pc, error) {
	c := vm.Pop().Data().(*gapl.Cont)
	return c.Resume(vm), nil
}

func (self *Resume) String() string {
	return fmt.Sprintf("RESUME")
}


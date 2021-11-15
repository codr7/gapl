package ops

import (
	"fmt"
	"github.com/codr7/gapl"
)

type Closure struct {
	form gapl.Form
	target *gapl.Func
}

func NewClosure(form gapl.Form, target *gapl.Func) *Closure {
	return &Closure{form: form, target: target}
}

func (self *Closure) Eval(pc gapl.Pc, vm *gapl.Vm) (gapl.Pc, error) {
	self.target.CaptureState(vm)
	return pc+1, nil
}

func (self *Closure) String() string {
	return fmt.Sprintf("CLOSURE %v", self.target)
}

package ops

import (
	"fmt"
	"github.com/codr7/gapl"
)

type Push struct {
	form gapl.Form
	val gapl.Val
}

func NewPush(form gapl.Form, _type gapl.Type, data interface{}) *Push {
	return &Push{form: form, val: gapl.NewVal(_type, data)}
}

func (self Push) Eval(pc gapl.Pc, vm *gapl.Vm) (gapl.Pc, error) {
	vm.Push(self.val.Type(), self.val.Data())
	return pc+1, nil
}

func (self Push) String() string {
	return fmt.Sprintf("PUSH %v", self.val)
}

package ops

import (
	"github.com/codr7/gapl"
)

type Push struct {
	form gapl.Form
	val gapl.Val
}

func NewPush(form gapl.Form, _type gapl.Type, data interface{}) *Push {
	return &Push{form: form, val: gapl.NewVal(_type, data)}
}

func (self Push) Eval(pc gapl.PC, vm *gapl.VM) (gapl.PC, error) {
	vm.Push(self.val.Type(), self.val.Data())
	return pc+1, nil
}

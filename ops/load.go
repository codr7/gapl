package ops

import (
	"fmt"
	"github.com/codr7/gapl"
)

type Load struct {
	form gapl.Form
	reg gapl.Reg
	_type gapl.Type
}

func NewLoad(form gapl.Form, reg gapl.Reg, _type gapl.Type) *Load {
	return &Load{form: form, reg: reg, _type: _type}
}

func (self Load) Eval(pc gapl.Pc, vm *gapl.Vm) (gapl.Pc, error) {
	vm.Load(self.reg)
	return pc+1, nil
}

func (self Load) String() string {
	return fmt.Sprintf("LOAD %v %v", self.reg, self._type)
}

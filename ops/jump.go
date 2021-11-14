package ops

import (
	"fmt"
	"github.com/codr7/gapl"
)

type Jump struct {
	form gapl.Form
	Pc gapl.Pc
}

func NewJump(form gapl.Form, pc gapl.Pc) *Jump {
	return &Jump{form: form, Pc: pc}
}

func (self *Jump) Eval(pc gapl.Pc, vm *gapl.Vm) (gapl.Pc, error) {
	return self.Pc, nil
}

func (self *Jump) String() string {
	return fmt.Sprintf("JUMP %v", self.Pc)
}


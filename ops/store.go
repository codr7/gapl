package ops

import (
	"fmt"
	"github.com/codr7/gapl"
)

type Store struct {
	form gapl.Form
	reg gapl.Reg
	_type gapl.Type
	Pop bool
}

func NewStore(form gapl.Form, reg gapl.Reg, _type gapl.Type) *Store {
	return &Store{form: form, reg: reg, _type: _type, Pop: true}
}

func (self *Store) Eval(pc gapl.Pc, vm *gapl.Vm) (gapl.Pc, error) {
	vm.Store(self.reg, self.Pop)
	return pc+1, nil
}

func (self *Store) Reg() gapl.Reg {
	return self.reg
}

func (self *Store) String() string {
	return fmt.Sprintf("STORE %v %v", self.reg, self._type)
}


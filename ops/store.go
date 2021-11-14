package ops

import (
	"fmt"
	"github.com/codr7/gapl"
)

type Store struct {
	form gapl.Form
	reg gapl.Reg
	_type gapl.Type
}

func NewStore(form gapl.Form, reg gapl.Reg, _type gapl.Type) *Store {
	return &Store{form: form, reg: reg, _type: _type}
}

func (self *Store) Eval(pc gapl.Pc, vm *gapl.Vm) (gapl.Pc, error) {
	vm.Store(self.reg)
	return pc+1, nil
}

func (self *Store) String() string {
	return fmt.Sprintf("STORE %v %v", self.reg, self._type)
}


package forms

import (
	"github.com/codr7/gapl"
)

type Group struct {
	gapl.BasicForm
	members []gapl.Form
}

func NewGroup(pos gapl.Pos, members...gapl.Form) *Group {
	self := new(Group)
	self.Init(pos)
	self.members = members
	return self
}

func (self *Group) Emit(in []gapl.Form, vm *gapl.Vm) ([]gapl.Form, error) {
	for _, m := range self.members {
		var err error

		if in, err = m.Emit(in, vm); err != nil {
			return in, err
		}
	}

	return in, nil
}

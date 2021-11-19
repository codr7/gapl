package forms

import (
	"github.com/codr7/gapl"
	"strings"
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

func (self *Group) Members() []gapl.Form {
	return self.members
}

func (self *Group) Emit(in []gapl.Form, vm *gapl.Vm) ([]gapl.Form, error) {
	temp := self.members[:]
	
	for len(temp) > 0 {
		var err error

		if temp, err = temp[0].Emit(temp[1:], vm); err != nil {
			return in, err
		}
	}

	return in, nil
}

func (self Group) String() string {
	var buf strings.Builder
	buf.WriteRune('(')

	for i, m := range self.members {
		if i > 0 {
			buf.WriteRune(' ')
		}
		
		buf.WriteString(m.String())
	}
	
	buf.WriteRune(')')
	return buf.String()
}

package forms

import (
	"github.com/codr7/gapl"
	"github.com/codr7/gapl/ops"
	"strings"
)

type Slice struct {
	gapl.BasicForm
	items []gapl.Form
	val *gapl.Val
}

func NewSlice(pos gapl.Pos, items...gapl.Form) *Slice {
	self := new(Slice)
	self.Init(pos)
	self.items = items
	return self
}

func (self *Slice) Items() []gapl.Form {
	return self.items
}

func (self *Slice) Emit(in []gapl.Form, vm *gapl.Vm) ([]gapl.Form, error) {
	temp := self.items[:]
	
	for len(temp) > 0 {
		var err error

		if temp, err = temp[0].Emit(temp[1:], vm); err != nil {
			return in, err
		}
	}

	vm.Emit(ops.NewSlice(self, len(self.items)))
	return in, nil
}

func (self Slice) String() string {
	var buf strings.Builder
	buf.WriteRune('[')

	for i, m := range self.items {
		if i > 0 {
			buf.WriteRune(' ')
		}
		
		buf.WriteString(m.String())
	}
	
	buf.WriteRune(']')
	return buf.String()
}

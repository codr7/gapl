package forms

import (
	"github.com/codr7/gapl"
)

type Literal struct {
	gapl.BasicForm
	val gapl.Val
}

func NewLiteral(pos gapl.Pos, _type gapl.Type, data interface{}) *Literal {
	self := new(Literal)
	self.Init(pos)
	self.val = gapl.NewVal(_type, data)
	return self
}

func (self *Literal) Emit(in []gapl.Form, vm *gapl.Vm) ([]gapl.Form, error) {
	return self.val.Emit(self, in, vm)
}

func (self Literal) Val(vm *gapl.Vm) *gapl.Val {
	return &self.val
}

func (self Literal) String() string {
	return self.val.String()
}

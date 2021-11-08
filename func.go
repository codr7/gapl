package gapl

import (
	"strings"
)

type Arg struct {
	name string
	_type Type
}

func NewArg(name string, _type Type) Arg {
	return Arg{name: name, _type: _type}
}

type Args []Arg

func (self Args) Add(name string, _type Type) Args {
	return append(self, NewArg(name, _type))
}

type Rets []Type

func (self Rets) Add(_type Type) Rets {
	return append(self, _type)
}

type CallFlags struct {
	Check bool
	Drop bool
	TCO bool
}

type FuncBody = func(self *Func, flags CallFlags, pc PC, vm *VM) (PC, error)

type Func struct {
	name string
	args Args
	rets Rets
	body FuncBody
}

func (self *Func) Init(name string, args Args, rets Rets, body FuncBody) *Func {
	self.name = name
	self.args = args
	self.rets = rets
	self.body = body
	return self
}

func (self *Func) Name() string {
	return self.name
}

func (self *Func) Dump() string {
	var buf strings.Builder
	buf.WriteString("Func(")
	buf.WriteString(self.name)
	buf.WriteString(" (")
	buf.WriteString("))")
	return buf.String()
}

func (self *Func) Call(flags CallFlags, pc PC, vm *VM) (PC, error) {
	return self.body(self, flags, pc, vm)
}

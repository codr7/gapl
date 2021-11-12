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

func (self Arg) Name() string {
	return self.name
}

func (self Arg) Type() Type {
	return self._type
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
	Tco bool
}

type FuncBody = func(self *Func, flags CallFlags, pc Pc, vm *Vm) (Pc, error)

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

func (self *Func) Args() Args {
	return self.args
}

func (self *Func) Rets() Rets {
	return self.rets
}

func (self *Func) String() string {
	var buf strings.Builder
	buf.WriteString("Func(")
	buf.WriteString(self.name)
	buf.WriteString(" [")
	buf.WriteString("] [")
	buf.WriteString("])")
	return buf.String()
}

func (self *Func) Applicable(stack *Stack) bool {
	nargs := len(self.args)
	its := stack.Items()
	nits := len(its)
	
	if nits < nargs {
		return false
	}
	
	for i := 0; i < nargs; i++ {
		if !Isa(its[nits-i-1].Type(), self.args[nargs-i-1]._type) {
			return false
		}
	}

	return true
}

func (self *Func) Call(flags CallFlags, pc Pc, vm *Vm) (Pc, error) {
	return self.body(self, flags, pc, vm)
}

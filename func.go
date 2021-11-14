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
	CheckArgs bool
	CheckRets bool
	Drop bool
	Tco bool
	Unsafe bool
}

type FuncBody = func(self *Func, flags CallFlags, retPc Pc, vm *Vm) (Pc, error)

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

func (self *Func) CompileBody(startPc Pc) {	
	self.body = func(self *Func, flags CallFlags, retPc Pc, vm *Vm) (Pc, error) {
		if !flags.Tco {
			f := vm.NewFrame(self, flags, retPc)
			f.CaptureState(vm)
		}
		
		return startPc, nil
	}
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
	if self.name != "" {
		buf.WriteString(self.name)
		buf.WriteRune(' ')
	}
	
	buf.WriteRune('(')

	for i, a := range self.args {
		if i > 0 {
			buf.WriteRune(' ')
		}
		
		buf.WriteString(a.name)
		buf.WriteRune(' ')
		buf.WriteString(a._type.Name())
	}
	
	buf.WriteString(") (")

	for i, rt := range self.rets {
		if i > 0 {
			buf.WriteRune(' ')
		}
		
		buf.WriteString(rt.Name())
	}

	buf.WriteString("))")
	return buf.String()
}

func (self *Func) Applicable(stack []Val) bool {
	nargs := len(self.args)
	nvals := len(stack)
	
	if nvals < nargs {
		return false
	}
	
	for i := 0; i < nargs; i++ {
		if !Isa(stack[nvals-i-1].Type(), self.args[nargs-i-1]._type) {
			return false
		}
	}

	return true
}

func (self *Func) Call(flags CallFlags, retPc Pc, vm *Vm) (Pc, error) {
	return self.body(self, flags, retPc, vm)
}

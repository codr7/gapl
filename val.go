package gapl

type Val struct {
	_type Type
	data interface{}
}

func NewVal(_type Type, data interface{}) Val {
	return Val{_type: _type, data: data}
}

func (self Val) Type() Type {
	return self._type
}

func (self Val) Data() interface{} {
	return self.data
}

func (self *Val) Set(_type Type, data interface{}) {
	self._type = _type
	self.data = data
}

func (self Val) String() string {
	return self._type.DumpVal(self)
}

func (self Val) Emit(form Form, in []Form, vm *Vm) ([]Form, error) {
	return self._type.EmitVal(self, form, in, vm)
}

func (self Val) Equal(other Val) bool {
	return self._type == other._type && self._type.EqualVals(self, other)
}

func (self Val) Literal() *Val {
	return self._type.LiteralVal(self)
}

func (self Val) True() bool {
	return self._type.TrueVal(self)
}

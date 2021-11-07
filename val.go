package gapl

type Val struct {
	_type Type
	data interface{}
}

func Val(_type Type, data interface{}) Val {
	return Val{_type: _type, data: data}
}

func (self Val) Type() Type {
	return self._type
}

func (self Val) Data() interface{} {
	return self.data
}

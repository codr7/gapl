package gapl

import (
	"fmt"
)

type Lib struct {
	name string
	bindings map[string]Val
}

func (self *Lib) Init(name string) *Lib {
	self.name = name
	self.bindings = make(map[string]Val)
	return self
}

func (self *Lib) Name() string {
	return self.name
}

func (self *Lib) Bind(key string, _type Type, data interface{}) error {
	prev, ok := self.bindings[key]

	if ok {
		return fmt.Errorf("Dup binding: %v\n%v", key, prev)
	}
	
	self.bindings[key] = NewVal(_type, data)
	return nil
}

func (self Lib) Find(key string) Val {
	return self.bindings[key]
}

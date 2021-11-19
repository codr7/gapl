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

func (self *Lib) Find(key string) *Val {
	if v, ok := self.bindings[key]; ok {
		return &v
	}

	return nil
}

func (self *Lib) Keys() []string {
	var out []string

	for k, _ := range self.bindings {
		out = append(out, k)
	}

	return out
}

func (self *Lib) Import(scope *Scope, keys...string) error {
	if keys == nil {
		keys = self.Keys()
	}
	
	for _, k := range keys {
		val := self.Find(k)

		if val == nil {
			return fmt.Errorf("Unknown id: %v", k)
		}

		if v := scope.Find(k); v != nil {
			return fmt.Errorf("Dup binding: %v\n%v", k, v)
		}
		
		scope.Bind(k, val.Type(), val.Data())
	}

	return nil
}

func (self Lib) String() string {
	return fmt.Sprintf("Lib(%v)", self.name)
}

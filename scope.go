package gapl

import (
)

type Scope struct {
	parentScope *Scope
	regCount int
	bindings map[string]Val
}

func (self *Scope) Init(parentScope *Scope) *Scope {
	self.parentScope = parentScope

	if parentScope != nil {
		self.regCount = parentScope.regCount
	}
	
	self.bindings = make(map[string]Val)
	return self
}

func (self *Scope) Bind(key string, _type Type, data interface{}) {
	self.bindings[key] = NewVal(_type, data)
}

func (self *Scope) Find(key string) *Val {
	v, ok := self.bindings[key]

	if !ok {
		if self.parentScope != nil {
			return self.parentScope.Find(key)
		}
		
		return nil
	}

	return &v
}

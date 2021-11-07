package gapl

type Scope struct {
	parentScope *Scope
	regCount int
	bindings map[string]Val
}

func (self *Scope) Init(parentScope *Scope) *Scope {
	self.parentScope = parentScope
	self.bindings = make(map[string]Val)
	return self
}

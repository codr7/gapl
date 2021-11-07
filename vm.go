package gapl

type PC int64

type VM struct {
	scope *Scope
	code []Op
	states []State
}

func (self *VM) BeginScope() {
	self.scope = new(Scope).Init(self.scope)
}

func (self *VM) EndScope() {
	self.scope = self.scope.parentScope
}

func (self *VM) BeginState() {
	self.states = append(self.states, State{})
}

func (self *VM) EndState() {
	self.states = self.states[:len(self.states)-1];
}

func (self *VM) State() *State {
	return &self.states[len(self.states)-1]
}

func (self *VM) Push(val Val) {
	self.State().Stack.Push(val)
}

func (self *VM) Peek() *Val {
	return self.State().Stack.Peek()
}

func (self *VM) PushNew(_type Type, data interface{}) {
	self.Push(NewVal(_type, data))
}

func (self *VM) Eval(pc PC) {
	for pc != -1 {
		pc = self.code[pc].Eval(pc, self)
	}
}

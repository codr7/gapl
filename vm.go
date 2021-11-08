package gapl

type PC int
type Reg int

type VM struct {
	scope *Scope
	code []Op
	states []State
}

func (self *VM) NewScope() {
	self.scope = new(Scope).Init(self.scope)
}

func (self *VM) EndScope() {
	self.scope = self.scope.parentScope
}

func (self *VM) Scope() *Scope {
	return self.scope
}

func (self *VM) Emit(op Op) Op {
	self.code = append(self.code, op)
	return op
}

func (self *VM) NewState() {
	self.states = append(self.states, State{})
}

func (self *VM) EndState() {
	self.states = self.states[:len(self.states)-1];
}

func (self *VM) State() *State {
	return &self.states[len(self.states)-1]
}

func (self *VM) Push(_type Type, data interface{}) {
	self.State().Stack.Push(_type, data)
}

func (self *VM) Peek() *Val {
	return self.State().Stack.Peek()
}

func (self *VM) Pop() Val {
	return self.State().Stack.Pop()
}

func (self *VM) Eval(pc PC) error {
	var err error
	
	for err == nil {
		pc, err = self.code[pc].Eval(pc, self)
	}

	if err != nil && err.Error() == "STOP" {
		return nil
	}

	return err
}

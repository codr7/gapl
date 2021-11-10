package gapl

import (
	"fmt"
)

type Pc int
type Reg int

func (self Reg) String() string { return fmt.Sprintf("Reg(%v)", self) }

type Vm struct {
	RegType Type
	scope *Scope
	code []Op
	states []State
}

func (self *Vm) NewScope() {
	self.scope = new(Scope).Init(self.scope)
}

func (self *Vm) EndScope() {
	self.scope = self.scope.parentScope
}

func (self *Vm) Scope() *Scope {
	return self.scope
}

func (self *Vm) Emit(op Op) Op {
	self.code = append(self.code, op)
	return op
}

func (self *Vm) NewState() {
	self.states = append(self.states, State{})
}

func (self *Vm) EndState() {
	self.states = self.states[:len(self.states)-1];
}

func (self *Vm) State() *State {
	return &self.states[len(self.states)-1]
}

func (self *Vm) Stack() *Stack {
	return &self.State().Stack
}

func (self *Vm) Push(_type Type, data interface{}) {
	self.State().Stack.Push(_type, data)
}

func (self *Vm) Peek() *Val {
	return self.State().Stack.Peek()
}

func (self *Vm) Pop() Val {
	return self.State().Stack.Pop()
}

func (self *Vm) Pc() Pc {
	return Pc(len(self.code))
}

func (self *Vm) Eval(pc Pc) error {
	var err error
	
	for err == nil {
		pc, err = self.code[pc].Eval(pc, self)
	}

	if err != nil && err.Error() == "STOP" {
		return nil
	}

	return err
}

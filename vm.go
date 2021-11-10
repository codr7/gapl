package gapl

import (
	"bufio"
	"fmt"
)

const VERSION = 1

type Pc int
type Reg int

func (self Reg) String() string { return fmt.Sprintf("Reg(%v)", self) }

type Vm struct {
	Readers []Reader
	IntType, RegType Type
	
	scope *Scope
	code []Op
	states []State
}

func (self *Vm) AddReader(in...Reader) {
	self.Readers = append(self.Readers, in...)
}

func (self *Vm) NewScope() {
	self.scope = new(Scope).Init(self.scope)
}

func (self *Vm) EndScope() {
	self.scope = self.scope.parentScope
}

func (self *Vm) ReadForm(in *bufio.Reader, pos *Pos) (Form, error) {
	for _, r := range self.Readers {
		if f, err := r(in, pos, self); f != nil || err != nil {
			return f, err
		}
	}
	
	return nil, nil
}

func (self *Vm) Scope() *Scope {
	return self.scope
}

func (self *Vm) Pc() Pc {
	return Pc(len(self.code))
}

func (self *Vm) Emit(op Op) {
	self.code = append(self.code, op)
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

func (self *Vm) BindReg(key string) int {
	scope := self.Scope()
	reg := scope.regCount
	scope.regCount++
	scope.Bind(key, self.RegType, reg)
	return reg
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

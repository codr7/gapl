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
	frame *Frame
	state *State
	code []Op
}

func (self *Vm) AddReader(in...Reader) {
	self.Readers = append(self.Readers, in...)
}

func (self *Vm) NewScope() *Scope {
	self.scope = new(Scope).Init(self.scope)
	return self.scope
}

func (self *Vm) EndScope() *Scope {
	s := self.scope
	self.scope = self.scope.parentScope
	return s
}

func (self *Vm) Scope() *Scope {
	return self.scope
}

func (self *Vm) ReadForm(in *bufio.Reader, pos *Pos) (Form, error) {
	for _, r := range self.Readers {
		if f, err := r(in, pos, self); f != nil || err != nil {
			return f, err
		}
	}
	
	return nil, nil
}

func (self *Vm) NewFrame(target *Func, flags CallFlags, retPc Pc) *Frame {
	self.frame = new(Frame).Init(self.frame, target, flags, retPc)
	self.frame.CaptureState(self)
	return self.frame
}

func (self *Vm) EndFrame() *Frame {
	f := self.frame
	self.frame = self.frame.parentFrame
	f.RestoreState(self)
	return f
}

func (self *Vm) Frame() *Frame {
	return self.frame
}

func (self *Vm) Pc() Pc {
	return Pc(len(self.code))
}

func (self *Vm) Emit(op Op) Op{
	self.code = append(self.code, op)
	return op
}

func (self *Vm) NewState() *State {
	self.state = new(State).Init(self.state)
	return self.state
}

func (self *Vm) EndState() *State {
	s := self.state
	self.state = self.state.parentState
	return s
}

func (self *Vm) State() *State {
	return self.state
}

func (self *Vm) Regs() []Val {
	return self.State().regs[:]
}

func (self *Vm) Stack() *Stack {
	return &self.State().stack
}

func (self *Vm) Push(_type Type, data interface{}) {
	self.State().stack.Push(_type, data)
}

func (self *Vm) Peek() *Val {
	return self.State().stack.Peek()
}

func (self *Vm) Pop() Val {
	return self.State().stack.Pop()
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

package gapl

import (
	"bufio"
	"fmt"
)

const VERSION = 1

type Pc int
type Reg int

func (self Reg) String() string { return fmt.Sprintf("Reg(%v)", int(self)) }

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
	return self.frame
}

func (self *Vm) EndFrame() *Frame {
	f := self.frame
	self.frame = self.frame.parentFrame
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

func (self *Vm) Load(reg Reg) {
	state := self.State()
	val := state.regs[reg]
	state.stack.Push(val.Type(), val.Data())
}

func (self *Vm) Store(reg Reg) {
	state := self.State()
	val := state.stack.Pop()
	state.regs[reg] = val
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

func (self *Vm) BindReg(key string) Reg {
	scope := self.Scope()

	if found := scope.Find(key); found != nil {
		return found.Data().(Reg)
	}
	
	reg := Reg(scope.regCount)
	scope.regCount++
	scope.Bind(key, self.RegType, reg)
	return reg
}

func (self *Vm) Eval(pc Pc) error {
	var err error
	
	for err == nil {
		//fmt.Printf("%v %v\n", self.code[pc], *self.Stack())
		pc, err = self.code[pc].Eval(pc, self)
	}

	if err != nil && err.Error() == "STOP" {
		return nil
	}

	return err
}

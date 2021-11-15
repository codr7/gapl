package gapl

import (
	"bufio"
	"fmt"
)

const VERSION = 4
const FRAME_COUNT = 64
const STATE_COUNT = 64

type Pc int
type Reg int
type Frames [FRAME_COUNT]Frame
type States [STATE_COUNT]State

func (self Reg) String() string { return fmt.Sprintf("Reg(%v)", int(self)) }

type Vm struct {
	Readers []Reader
	BoolType, ContType, IntType, RegType Type
	
	scope *Scope
	frames Frames
	frameCount int
	states States
	stateCount int
	code []Op
	unsafeDepth int
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
	self.scope = s.parentScope
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
	if self.frameCount == FRAME_COUNT {
		panic("No more frames!")
	}
	
	f := &self.frames[self.frameCount]
	f.Init(target, flags, retPc)
	self.frameCount++
	return f 
}

func (self *Vm) EndFrame() *Frame {
	self.frameCount--
	return &self.frames[self.frameCount]
}

func (self *Vm) Frame() *Frame {
	return &self.frames[self.frameCount-1]
}

func (self *Vm) Pc() Pc {
	return Pc(len(self.code))
}

func (self *Vm) Emit(op Op) Op{
	//fmt.Printf(":%v\n", op)
	self.code = append(self.code, op)
	return op
}

func (self *Vm) NewState() *State {
	if self.stateCount == STATE_COUNT {
		panic("No more states!")
	}
	
	s := &self.states[self.stateCount]
	s.Init()
	self.stateCount++
	return s
}

func (self *Vm) EndState() *State {
	self.stateCount--
	return &self.states[self.stateCount]
}

func (self *Vm) State() *State {
	return &self.states[self.stateCount-1]
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

func (self *Vm) Unsafe() bool {
	return self.unsafeDepth > 0
}

package gapl

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const VERSION = 7
const FRAME_COUNT = 64
const STATE_COUNT = 64

type Pc int
type Reg int
type Frames [FRAME_COUNT]Frame
type States [STATE_COUNT]State

func (self Reg) String() string { return fmt.Sprintf("Reg(%v)", int(self)) }

type Vm struct {
	Readers []Reader
	BoolType, ContType, IntType, RegType, StringType Type
	
	scope *Scope
	frames Frames
	frameCount int
	states States
	stateCount int
	code []Op

	path string
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

func (self *Vm) Bind(key string, _type Type, data interface{}) {
	self.Scope().Bind(key, _type, data)
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

func (self *Vm) Find(key string) *Val {
	return self.Scope().Find(key)
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

func (self *Vm) Include(path string) error {
	path = filepath.Join(self.path, path)
	f, err := os.Open(path)

	if err != nil {
		return fmt.Errorf("Failed opening file: %v %v", path, err)
	}
	
	bin := bufio.NewReader(f)
	pos := NewPos(path, 0, 0)
	var forms []Form

	for {
		if f, err := self.ReadForm(bin, &pos); err == io.EOF {
			break
		} else if err != nil {
			return err
		} else if f == nil {
			break
		} else {
			forms = append(forms, f)
		}
	}

	for len(forms) != 0 {
		f := forms[0]
		var err error
		forms, err = f.Emit(forms[1:], self)
		
		if err != nil {
			return err
		}
	}
	
	self.Emit(&STOP)
	return nil
}

func (self *Vm) Repl(in io.Reader, out io.Writer) {
	fmt.Fprintf(out, "  ")
	var buf strings.Builder
	ins := bufio.NewScanner(in)
	
	for ins.Scan() {
		if line := ins.Text(); len(line) == 0 && buf.Len() > 0 {
			bin := bufio.NewReader(strings.NewReader(buf.String()))
			pos := NewPos("repl", 0, 0)
			var forms []Form
			
			for {
				if f, err := self.ReadForm(bin, &pos); err == io.EOF {
					break
				} else if err != nil {
					fmt.Fprintln(out, err)
					forms = nil
					break
				} else if f == nil {
					break
				} else {
					forms = append(forms, f)
				}
			}

			pc := self.Pc()
			
			for len(forms) != 0 {
				f := forms[0]
				var err error
				forms, err = f.Emit(forms[1:], self)
				
				if err != nil {
					fmt.Fprintln(out, err)
					break
				}
			}

			if len(forms) == 0 && self.Pc() != pc {
				self.Emit(&STOP)
				
				if err := self.Eval(pc); err != nil {
					fmt.Fprintln(out, err)
				}
			}
			
			buf.Reset()
			fmt.Fprintf(out, "%v\n", *self.Stack())
		} else {
			buf.WriteString(line)
			buf.WriteRune('\n')
		}

		fmt.Fprintf(out, "  ")
	}
}

func (self *Vm) Unsafe() bool {
	return self.unsafeDepth > 0
}

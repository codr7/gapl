package core

type PC Int

type VM struct {
	code []Op
	states []State
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
	self.State().Push(val)
}

func (self *VM) Eval(pc PC) {
	for pc != -1 {
		pc = self.code[pc].Eval(pc, self)
	}
}

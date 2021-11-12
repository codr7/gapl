package gapl

const REG_COUNT = 64

type State struct {
	parentState *State
	regs [REG_COUNT]Val
	stack Stack
}

func (self *State) Init(parentState *State) *State {
	self.parentState = parentState
	return self
}

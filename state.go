package gapl

const REG_COUNT = 64

type Regs [REG_COUNT]Val

type State struct {
	regs Regs
	stack Stack
}

func (self *State) Stack() *Stack {
	return &self.stack
}

package gapl

const REG_COUNT = 64

type State struct {
	regs [REG_COUNT]Val
	stack Stack
}

func (self State) Stack() *Stack {
	return &self.stack
}

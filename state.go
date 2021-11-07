package gapl

const REG_COUNT = 64

type State struct {
	Regs [REG_COUNT]Val
	Stack Stack
}

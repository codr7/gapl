package core

type State struct {
	stack []Val
}

func (self *State) Push(val Val) {
	self.stack = append(self.stack, val)
}

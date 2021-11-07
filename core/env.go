package core

type Env struct {
	stack []Val
}

func (self *Env) Push(val Val) {
	self.stack = append(self.stack, val)
}

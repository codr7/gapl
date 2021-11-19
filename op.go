package gapl

import (
	"fmt"
)

type Op interface {
	Eval(pc Pc, vm *Vm) (Pc, error)
}

type EEval struct {
	E
	pos Pos
}

func NewEEval(pos Pos, message string, args...interface{}) EEval {
	var self EEval
	self.Init(message, args...)
	self.pos = pos
	return self
}

func (self EEval) Error() string {
	return fmt.Sprintf("Error in %v: %v", self.pos, self.message)
}

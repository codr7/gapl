package gapl

import (
	"bufio"
	"fmt"
)

type ERead struct {
	E
	pos Pos
}

func NewERead(pos Pos, message string, args...interface{}) ERead {
	var self ERead
	self.Init(message, args...)
	self.pos = pos
	return self
}

func (self ERead) Error() string {
	return fmt.Sprintf("Error in %v: %v", self.pos, self.message)
}

type Reader func(in *bufio.Reader, pos *Pos, vm *Vm) (Form, error)

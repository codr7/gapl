package gapl

import (
	"fmt"
)

type EEmit struct {
	E
	pos Pos
}

func NewEEmit(pos Pos, message string, args...interface{}) EEmit {
	var self EEmit
	self.Init(message, args...)
	self.pos = pos
	return self
}

func (self EEmit) Error() string {
	return fmt.Sprintf("Error in %v: %v", self.pos, self.message)
}

type Form interface {
	Pos() Pos
	Emit(in []Form, vm *Vm) ([]Form, error)
	Val(vm *Vm) *Val
}

type BasicForm struct {
	pos Pos
}

func (self *BasicForm) Init(pos Pos) *BasicForm {
	self.pos = pos
	return self
}

func (self BasicForm) Pos() Pos {
	return self.pos
}

func (self BasicForm) Val(vm *Vm) *Val {
	return nil
}

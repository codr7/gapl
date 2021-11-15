package gapl

import (
	"fmt"
)

type Cont struct {
	states States
	stateCount int
	
	frames Frames
	frameCount int
	
	pc Pc
}

func (self *Cont) Init(pc Pc) *Cont {
	self.pc = pc
	return self
}

func (self *Cont) Suspend(vm *Vm) {
	self.frames = vm.frames
	self.frameCount = vm.frameCount

	self.states = vm.states
	self.stateCount = vm.stateCount

	vm.frameCount = 0
	vm.stateCount = 0
}

func (self *Cont) Resume(vm *Vm) Pc {
	vm.frames = self.frames
	vm.frameCount = self.frameCount

	vm.states = self.states
	vm.stateCount = self.stateCount
	return self.pc
}

func (self *Cont) String() string {
	return fmt.Sprintf("Cont(%v)", self.pc)
}


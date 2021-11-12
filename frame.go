package gapl

import (
	"fmt"
)

type Frame struct {
	parentFrame *Frame
	target *Func
	flags CallFlags
	retPc Pc
}

func (self *Frame) Init(parentFrame *Frame, target *Func, flags CallFlags, retPc Pc) *Frame {
	self.parentFrame = parentFrame
	self.target = target
	self.flags = flags
	self.retPc = retPc
	return self
}

func (self *Frame) CaptureState(vm *Vm) {
	src, dst := vm.State(), vm.NewState()
	
	if args := self.target.Args(); len(args) > 0 {
		dst.stack.Append(src.stack.Items()[src.stack.Len()-len(args):])
		src.stack.Drop(len(args))
	}

	dst.regs = src.regs
}

func (self *Frame) RestoreState(vm *Vm) error {
	src, dst := vm.EndState(), vm.State()
	rets := self.target.Rets()
		
	if len(rets) > 0 {
		if src.stack.Len() < len(rets) {
			return fmt.Errorf("Not enough return values: %v\n%v", len(rets), src.stack)
		}
			
		for i, rt := range rets {
			st := src.stack.Items()[src.stack.Len()-i-1].Type()
			
			if !Isa(st, rt) {
				return fmt.Errorf("Wrong return type: %v %v", st, rt)
			}
		}

		
		if !self.flags.Drop {
			dst.stack.Append(src.stack.Items()[src.stack.Len()-len(rets):])
		}
	}

	return nil
}

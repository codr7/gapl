package gapl

import (
	"fmt"
)

type Frame struct {
	target *Func
	flags CallFlags
	retPc Pc
}

func (self *Frame) Init(target *Func, flags CallFlags, retPc Pc) *Frame {
	self.target = target
	self.flags = flags
	self.retPc = retPc
	return self
}

func (self *Frame) CaptureState(vm *Vm) {
	src, dst := vm.State(), vm.NewState()
	
	if argCount := len(self.target.Args()); argCount > 0 {
		dst.stack.Append(src.stack.Drop(argCount))
	}

	dst.regs = self.target.regs

	if self.flags.Unsafe {
		vm.unsafeDepth++
	}
}

func (self *Frame) RestoreState(vm *Vm) (Pc, error) {
	rets := self.target.Rets()
	src := vm.EndState()
	
	if retCount := len(rets); retCount > 0 {
		dst := vm.State()
		
		if self.flags.CheckRets && !vm.Unsafe() {
			valCount := src.stack.Len()
			
			if valCount < retCount {
				return -1, fmt.Errorf("Missing return values: %v %v", retCount, src.stack)
			}

			for i, rt := range rets {
				st := src.stack.Items()[valCount-i-1].Type()
				
				if !Isa(st, rt) {
					return -1, fmt.Errorf("Wrong type returned: %v %v", st, rt)
				}
			}
		}

		
		if !self.flags.Drop {
			dst.stack.Append(src.stack.Drop(retCount))
		}
	}

	if self.flags.Unsafe {
		vm.unsafeDepth--
	}
	
	return self.retPc, nil
}

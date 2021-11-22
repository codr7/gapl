package procs

import (
	"github.com/codr7/gapl"
	"github.com/codr7/gapl/ops"
)

func Fuse(pc gapl.Pc, vm *gapl.Vm) {
	var prev gapl.Op
	
	for i := int(pc); i < len(vm.Code); i++ {
		op := vm.Code[i]
		
		if prevStore, ok := prev.(*ops.Store); ok {
			if op, ok := op.(*ops.Load); ok && op.Reg() == prevStore.Reg() {
				prevStore.Pop = false
				vm.Code[i] = &gapl.NOP
				continue
			}
		}

		prev = op
	}
}

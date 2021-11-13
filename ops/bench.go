package ops

import (
	"fmt"
	"github.com/codr7/gapl"
	"time"
)

type Bench struct {
	form gapl.Form
	EndPc gapl.Pc
}

func NewBench(form gapl.Form, endPc gapl.Pc) *Bench {
	return &Bench{form: form, EndPc: endPc}
}

func (self Bench) Eval(pc gapl.Pc, vm *gapl.Vm) (gapl.Pc, error) {
	reps := vm.Pop().Data().(int)
	start := time.Now()

	for i := 0; i < reps; i++ {
		if err := vm.Eval(pc+1); err != nil {
			return -1, err
		}
	}
	
	ms := time.Since(start).Milliseconds()
	vm.Push(vm.IntType, int(ms))
	return self.EndPc, nil
}

func (self Bench) String() string {
	return fmt.Sprintf("BENCH %v", self.EndPc)
}

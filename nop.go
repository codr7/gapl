package gapl

type Nop struct {}

var NOP Nop

func (self *Nop) Eval(pc Pc, vm *Vm) (Pc, error) {
	return pc+1, nil
}

func (self *Nop) String() string {
	return "NOP"
}

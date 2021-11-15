package gapl

type Stop struct {}

var STOP Stop

func (self *Stop) Error() string {
	return "STOP"
}

func (self *Stop) Eval(pc Pc, vm *Vm) (Pc, error) {
	return pc, self
}

func (self *Stop) String() string {
	return "STOP"
}

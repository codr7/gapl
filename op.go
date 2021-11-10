package gapl

type Op interface {
	Eval(pc Pc, vm *Vm) (Pc, error)
}

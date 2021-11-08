package gapl

type Op interface {
	Eval(pc PC, vm *VM) (PC, error)
}

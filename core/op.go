package core

type Op interface {
	Eval(pc PC, vm *VM) PC
}

package ops

type Push struct {
	val core.Val
}

func (self Push) Eval(pc core.PC, vm *core.VM) PC {
	vm.Push(self.val)
	return pc+1
}

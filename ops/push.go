package ops

type Push struct {
	val gapl.Val
}

func (self Push) Eval(pc gapl.PC, vm *gapl.VM) gapl.PC {
	vm.Push(self.val)
	return pc+1
}

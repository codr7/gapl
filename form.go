package gapl

type Form interface {
	Pos() Pos
	Emit(in []Form, vm *VM) ([]Form, error)
	Val(vm *VM) *Val
}

type BasicForm struct {
	pos Pos
}

func (self *BasicForm) Init(pos Pos) *BasicForm {
	self.pos = pos
	return self
}

func (self BasicForm) Pos() Pos {
	return self.pos
}

func (self BasicForm) Val(vm *VM) *Val {
	return nil
}

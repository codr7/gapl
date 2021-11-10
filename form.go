package gapl

type Form interface {
	Pos() Pos
	Emit(in []Form, vm *Vm) ([]Form, error)
	Val(vm *Vm) *Val
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

func (self BasicForm) Val(vm *Vm) *Val {
	return nil
}

package gapl

import (
	"fmt"
)

type Type interface {
	Init(name string, parentTypes...Type)

	Name() string
	AddParentTypes(dst map[Type]Type, dpt Type)
	GetParentType(other Type) Type

	DumpVal(v Val) string
	EmitVal(v Val, form Form, in []Form, vm *Vm) ([]Form, error)
	EqualVals(x, y Val) bool
	TrueVal(v Val) bool
}

type BasicType struct {
	name string
	parentTypes map[Type]Type
}

func (self *BasicType) Init(name string, parentTypes...Type) {
	self.name = name
	self.parentTypes = make(map[Type]Type)
	
	for _, pt := range parentTypes {
		self.Derive(pt)
	}
}

func (self BasicType) Name() string {
	return self.name
}

func (self BasicType) String() string {
	return self.Name()
}

func (self BasicType) AddParentTypes(dst map[Type]Type, dpt Type) {
	for pt, _ := range self.parentTypes {
		dst[pt] = dpt
	}
}

func (self *BasicType) Derive(other Type) {
	self.parentTypes[other] = other
	other.AddParentTypes(self.parentTypes, other)
}

func (self BasicType) GetParentType(other Type) Type {
	return self.parentTypes[other]
}

func (self BasicType) DumpVal(v Val) string {
	return fmt.Sprintf("%v", v.Data())
}

func (self BasicType) EqualVals(x, y Val) bool {
	return x.Data() == y.Data()
}

func (self BasicType) TrueVal(v Val) bool {
	return true
}


func Isa(child, parent Type) bool {
	return child == parent || child.GetParentType(parent) != nil
}

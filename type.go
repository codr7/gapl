package gapl

import (
	"fmt"
)

const TYPE_COUNT = 64

var nextTypeId int

type Type interface {
	Init(name string, parentTypes...Type)

	Id() int
	Name() string
	AddParentTypes(dst [TYPE_COUNT]Type, dpt Type)
	GetParentType(other Type) Type

	DumpVal(v Val) string
	EmitVal(v Val, form Form, in []Form, vm *Vm) ([]Form, error)
	EqualVals(x, y Val) bool
	LiteralVal(v Val) *Val
	TrueVal(v Val) bool
}

type BasicType struct {
	id int
	name string
	parentTypes [TYPE_COUNT]Type
}

func (self *BasicType) Init(name string, parentTypes...Type) {
	self.id = nextTypeId
	nextTypeId++
	
	self.name = name
	
	for _, pt := range parentTypes {
		self.Derive(pt)
	}
}

func (self *BasicType) Id() int {
	return self.id
}

func (self *BasicType) Name() string {
	return self.name
}

func (self *BasicType) String() string {
	return self.Name()
}

func (self *BasicType) AddParentTypes(dst [TYPE_COUNT]Type, dpt Type) {
	for i, pt := range self.parentTypes {
		if pt != nil {
			dst[i] = dpt
		}
	}
}

func (self *BasicType) Derive(other Type) {
	self.parentTypes[other.Id()] = other
	other.AddParentTypes(self.parentTypes, other)
}

func (self *BasicType) GetParentType(other Type) Type {
	return self.parentTypes[other.Id()]
}

func (self *BasicType) DumpVal(v Val) string {
	return fmt.Sprintf("%v", v.Data())
}

func (self *BasicType) EqualVals(x, y Val) bool {
	return x.Data() == y.Data()
}

func (self *BasicType) LiteralVal(v Val) *Val {
	return &v
}

func (self *BasicType) TrueVal(v Val) bool {
	return true
}

func Isa(child, parent Type) bool {
	return child == parent || child.GetParentType(parent) != nil
}

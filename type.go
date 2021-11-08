package gapl

import (
	"fmt"
)

type Type interface {
	Name() string
	AddParentTypes(dst map[Type]Type, dpt Type)
	GetParentType(other Type) Type

	DumpVal(v interface{}) string
}

type BasicType struct {
	name string
	parentTypes map[Type]Type
}

func (self *BasicType) Init(name string, parentTypes...Type) *BasicType {
	self.name = name
	self.parentTypes = make(map[Type]Type)
	
	for _, pt := range parentTypes {
		self.Derive(pt)
	}
	
	return self
}

func (self *BasicType) Name() string {
	return self.name
}

func (self *BasicType) AddParentTypes(dst map[Type]Type, dpt Type) {
	for pt, _ := range self.parentTypes {
		dst[pt] = dpt
	}
}

func (self *BasicType) Derive(other Type) {
	self.parentTypes[other] = other
	other.AddParentTypes(self.parentTypes, other)
}

func (self *BasicType) GetParentType(other Type) Type {
	return self.parentTypes[other]
}

func (self *BasicType) DumpVal(v interface{}) string {
	return fmt.Sprintf("%v", v)
}

func Isa(child, parent Type) bool {
	return child == parent || child.GetParentType(parent) != nil
}

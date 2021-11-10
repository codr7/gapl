package gapl

import (
	"fmt"
)

type Pos struct {
	source string
	line, column int
}

func NewPos(source string, line, column int) Pos {
	return Pos{source: source, line: line, column: column}
}


func (self *Pos) Read(in rune) {
	switch in {
	case '\n':
		self.line++
		self.column = 0
	default:
		self.column++
	}
}

func (self Pos) String() string {
	return fmt.Sprintf("%v at line %v, column %v", self.source, self.line, self.column)
}

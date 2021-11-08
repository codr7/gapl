package gapl

type Pos struct {
	source string
	line, column int
}

func NewPos(source string, line, column int) Pos {
	return Pos{source: source, line: line, column: column}
}

func (self *Pos) NewLine() {
	self.line++
	self.column = 0
}

func (self *Pos) NextColumn() {
	self.column++
}

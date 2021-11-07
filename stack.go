package gapl

import (
	"strings"
)

type Stack struct {
	Items []Val
}

func (self *Stack) Push(it Val) {
	self.Items = append(self.Items, it)
}

func (self *Stack) Peek() *Val {
	return &self.Items[len(self.Items)-1]
}

func (self *Stack) Pop() Val {
	i := len(self.Items)-1
	it := self.Items[i]
	self.Items = self.Items[:i];
	return it
}

func (self Stack) Dump() string {
	var buf strings.Builder
	buf.WriteRune('[')

	for i := 0; i < len(self.Items); i++ {
		if i > 0 {
			buf.WriteRune(' ')
		}

		buf.WriteString(self.Items[i].Dump())
	}

	buf.WriteRune(']')
	return buf.String()
}


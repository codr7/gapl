package gapl

import (
	"strings"
)

type Stack struct {
	items []Val
}

func (self *Stack) Len() int {
	return len(self.items)
}

func (self *Stack) Items() []Val {
	return self.items
}

func (self *Stack) Append(items []Val) {
	self.items = append(self.items, items...)
}

func (self *Stack) Drop(count int) {
	self.items = self.items[:len(self.items)-count]
}

func (self *Stack) Push(_type Type, data interface{}) {
	self.items = append(self.items, NewVal(_type, data))
}

func (self *Stack) Peek() *Val {
	return &self.items[len(self.items)-1]
}

func (self *Stack) Pop() Val {
	i := len(self.items)-1
	it := self.items[i]
	self.items = self.items[:i]
	return it
}

func (self Stack) String() string {
	var buf strings.Builder
	buf.WriteRune('[')

	for i := 0; i < len(self.items); i++ {
		if i > 0 {
			buf.WriteRune(' ')
		}

		buf.WriteString(self.items[i].Dump())
	}

	buf.WriteRune(']')
	return buf.String()
}

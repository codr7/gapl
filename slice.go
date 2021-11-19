package gapl

import (
	"strings"
)

type Slice []Val

func (self Slice) String() string {
	var buf strings.Builder
	buf.WriteRune('[')

	for i, it := range self {
		if i > 0 {
			buf.WriteRune(' ')
		}

		buf.WriteString(it.Dump())
	}

	buf.WriteRune(']')
	return buf.String()
}


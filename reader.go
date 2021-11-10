package gapl

import (
	"bufio"
	"fmt"
	"unicode"
)

type ERead struct {
	E
	pos Pos
}

func NewERead(pos Pos, message string, args...interface{}) ERead {
	var self ERead
	self.Init(message, args...)
	self.pos = pos
	return self
}

func (self ERead) Error() string {
	return fmt.Sprintf("Error in %v: %v", self.pos, self.message)
}

type Reader func(in *bufio.Reader, pos *Pos, vm *VM) (Form, error)

var AllReaders = []Reader{ReadWs, ReadId}

func ReadForm(in *bufio.Reader, pos *Pos, vm *VM) (Form, error) {
	for _, r := range AllReaders {
		if f, err := r(in, pos, vm); f != nil || err != nil {
			return f, err
		}
	}
	
	return nil, nil
}

func ReadId(in *bufio.Reader, pos *Pos, vm *VM) (Form, error) {
	fpos := *pos
	//var name strings.Builder
	return nil, NewERead(fpos, "Invalid input")
}

func ReadWs(in *bufio.Reader, pos *Pos, vm *VM) (Form, error) {
	for {
		if c, _, err := in.ReadRune(); err != nil {
			return nil, err
		} else if unicode.IsSpace(c) {
			pos.Read(c)
		} else {
			in.UnreadRune()
			break
		}
	}
	
	return nil, nil
}

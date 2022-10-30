package utils

import "fmt"

type ParseError struct {
	Pos  int
	msg  string
	args []string
}

func NewParseError(pos int, msg string, args []string) *ParseError {
	return &ParseError{pos, msg, args}
}

func (p ParseError) Error() string {
	return fmt.Sprintf("%s at position %d", fmt.Sprintf(p.msg, p.args), p.Pos)
}

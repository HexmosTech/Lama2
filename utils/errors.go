package utils

import "fmt"

type ParseError struct {
	Pos     int
	LineNum int
	msg     string
	args    []interface{}
}

func NewParseError(pos int, line int, msg string, args []string) *ParseError {
	argsInt := make([]interface{}, 0)
	for _, e := range args {
		argsInt = append(argsInt, e)
	}
	return &ParseError{pos, line, msg, argsInt}
}

func (p ParseError) Error() string {
	return fmt.Sprintf("%s at position %d, line %d", fmt.Sprintf(p.msg, p.args...), p.Pos, p.LineNum)
}

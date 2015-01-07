package parser

type SyntaxError interface {
	error
	Line() int
	Column() int
}

type syntaxError struct {
	line int
	col  int
	msg  string
}

func NewSyntaxError(message string, line, column int) SyntaxError {
	return syntaxError{
		msg:  message,
		line: line,
		col:  column,
	}
}

func (se syntaxError) Line() int {
	return se.line
}

func (se syntaxError) Column() int {
	return se.col
}

func (se syntaxError) Error() string {
	return se.msg
}

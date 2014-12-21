package statementTypes

import (
	"errors"
	"fmt"
	"github.com/apoydence/GoF/parser"
)

type ReturnStatement struct {
	block       string
	outputQueue []string
}

func NewReturnStatementParser() StatementParser {
	return ReturnStatement{}
}

func newReturnStatement(block string) Statement {
	statement := &ReturnStatement{
		block: block,
	}

	return statement
}

func (rs *ReturnStatement) OutputQueue() []string {
	return rs.outputQueue
}

func (rs ReturnStatement) Parse(block string, nextBlockScanner *parser.ScanPeeker, factory *StatementFactory) Statement {
	return newReturnStatement(block)
}

func (ds *ReturnStatement) GenerateGo(fm FunctionMap) (string, error) {
	ops, err := toRpn(ds.block, []string{}, []string{})
	if err != nil {
		return "", err
	}

	return wrapCode(toInfix(ops, 0))
}

func wrapCode(code string, err error) (string, error) {
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("func() int {\n\t%s\n}", code), nil
}

// This uses the shunting-yard algorithm
func toRpn(line string, outputQueue []string, opStack []string) ([]string, error) {
	token, rest := parser.Tokenize(line)
	if parser.IsPrimitive(token) {
		return toRpn(rest, append(outputQueue, token), opStack)
	} else if parser.IsOperator(token) {
		if len(opStack) > 0 {
			prec := opPrec(token)
			stackTop := opStack[len(opStack)-1]
			if parser.IsOperator(stackTop) && prec <= opPrec(stackTop) {
				opStack[len(opStack)-1] = token
				return toRpn(rest, append(outputQueue, stackTop), opStack)
			} else {
				return toRpn(rest, outputQueue, append(opStack, token))
			}
		} else {
			return toRpn(rest, outputQueue, append(opStack, token))
		}
	} else if token == "(" {
		return toRpn(rest, outputQueue, append(opStack, token))
	} else if token == ")" {
		var i int
		for i = len(opStack) - 1; opStack[i] != "("; i-- {
			outputQueue = append(outputQueue, opStack[i])
		}
		return toRpn(rest, outputQueue, opStack[:i])
	} else if token == "" {
		for i := len(opStack) - 1; i > -1; i-- {
			outputQueue = append(outputQueue, opStack[i])
		}
		return outputQueue, nil
	}

	return nil, errors.New("Unknown token: " + token)
}

func toInfix(opQueue []string, index int) (string, error) {
	if len(opQueue) <= index {
		return opQueue[0], nil
	} else if parser.IsOperator(opQueue[index]) {
		combined := "(" + opQueue[index-2] + opQueue[index] + opQueue[index-1] + ")"
		return toInfix(append(opQueue[:index-2], combine(combined, opQueue[index+1:])...), index-2)
	} else {
		return toInfix(opQueue, index+1)
	}
}

func combine(x string, y []string) []string {
	return append([]string{x}, y...)
}

func opPrec(op string) int {
	if op == "+" || op == "-" {
		return 0
	} else if op == "*" || op == "/" {
		return 1
	}
	panic("Unknown op: '" + op + "'")
}

func isFunc(a string) bool {
	return a != ")"
}

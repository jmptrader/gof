package expressionParsing

import (
	"errors"
	"github.com/apoydence/GoF/parser"
)

func toRpn(line string, outputQueue []rpnValue, opStack []rpnValue, fm FunctionMap) ([]string, error) {
	token, rest := parser.Tokenize(line)
	if parser.IsPrimitive(token) {
		return toRpn(rest, append(outputQueue, newPrimRpnValue(token)), opStack, fm)
	} else if parser.IsOperator(token) {
		return toRpnOperator(token, rest, outputQueue, opStack, fm)
	} else if token == "(" {
		return toRpn(rest, outputQueue, append(opStack, newParenRpnValue()), fm)
	} else if token == ")" {
		return toRpnLeftParen(token, rest, outputQueue, opStack, fm)
	} else if fd := fm.GetFunction(token); fd != nil {

	} else if token == "" {
		for i := len(opStack) - 1; i > -1; i-- {
			outputQueue = append(outputQueue, opStack[i])
		}
		return rpnToString(outputQueue), nil
	}

	return nil, errors.New("Unknown token: " + token)
}

func toRpnOperator(token, rest string, outputQueue []rpnValue, opStack []rpnValue, fm FunctionMap) ([]string, error) {
	prec := opPrec(token)
	op := newOpRpnValue(token, prec)
	if len(opStack) > 0 {
		stackTop := opStack[len(opStack)-1]
		if stackTop.operator && prec <= stackTop.prec {
			opStack[len(opStack)-1] = op
			return toRpn(rest, append(outputQueue, stackTop), opStack, fm)
		} else {
			return toRpn(rest, outputQueue, append(opStack, op), fm)
		}
	} else {
		return toRpn(rest, outputQueue, append(opStack, op), fm)
	}
}

func toRpnLeftParen(token, rest string, outputQueue []rpnValue, opStack []rpnValue, fm FunctionMap) ([]string, error) {
	var i int
	for i = len(opStack) - 1; !opStack[i].leftPar; i-- {
		outputQueue = append(outputQueue, opStack[i])
	}
	return toRpn(rest, outputQueue, opStack[:i], fm)
}

func rpnToString(ops []rpnValue) []string {
	result := make([]string, 0)
	for _, ops := range ops {
		result = append(result, ops.token)
	}

	return result
}

func opPrec(op string) int {
	if op == "+" || op == "-" {
		return 0
	} else if op == "*" || op == "/" {
		return 1
	}
	panic("Unknown op: '" + op + "'")
}

func combine(x *blockSpec, y []*blockSpec) []*blockSpec {
	return append([]*blockSpec{x}, y...)
}

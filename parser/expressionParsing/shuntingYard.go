package expressionParsing

import (
	"github.com/apoydence/gof/parser"
)

func toRpn(line string, outputQueue []rpnValue, opStack []rpnValue, fm FunctionMap) ([]string, parser.SyntaxError) {
	token, rest := parser.Tokenize(line)
	if parser.IsOperator(token) {
		return toRpnOperator(token, rest, outputQueue, opStack, fm)
	} else if token == "(" {
		return toRpn(rest, outputQueue, append(opStack, newParenRpnValue()), fm)
	} else if token == ")" {
		return toRpnRightParen(token, rest, outputQueue, opStack, fm)
	} else if fd := fm.GetFunction(token); fd != nil && fd.IsFunc() {
		ftd := fd.(FuncTypeDefinition)
		if ftd.IsDefinition() {
			return toRpn(rest, append(outputQueue, newPrimRpnValue(token)), opStack, fm)
		} else {
			return toRpn(rest, outputQueue, append(opStack, newOpRpnValue(token, FuncCall)), fm)
		}
	} else if parser.IsPrimitive(token) || parser.ValidFunctionName(token) {
		return toRpn(rest, append(outputQueue, newPrimRpnValue(token)), opStack, fm)
	} else if token == "" {
		for i := len(opStack) - 1; i > -1; i-- {
			outputQueue = append(outputQueue, opStack[i])
		}
		return rpnToString(outputQueue), nil
	}

	return nil, parser.NewSyntaxError("Unknown token: "+token, 0, 0)
}

func toRpnOperator(token, rest string, outputQueue []rpnValue, opStack []rpnValue, fm FunctionMap) ([]string, parser.SyntaxError) {
	prec := opPrec(token)
	op := newOpRpnValue(token, prec)
	if len(opStack) > 0 {
		stackTop := opStack[len(opStack)-1]
		if stackTop.operator && prec <= stackTop.prec {
			opStack[len(opStack)-1] = op
			return toRpnOperator(token, rest, append(outputQueue, stackTop), opStack[:len(opStack)-1], fm)
		} else {
			return toRpn(rest, outputQueue, append(opStack, op), fm)
		}
	} else {
		return toRpn(rest, outputQueue, append(opStack, op), fm)
	}
}

func toRpnRightParen(token, rest string, outputQueue []rpnValue, opStack []rpnValue, fm FunctionMap) ([]string, parser.SyntaxError) {
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
		return AddSub
	} else if op == "*" || op == "/" {
		return MultDiv
	}
	panic("Unknown op: '" + op + "'")
}

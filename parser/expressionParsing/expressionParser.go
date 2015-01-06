package expressionParsing

import "github.com/apoydence/GoF/parser"

type FunctionMap interface {
	GetFunction(name string) TypeDefinition
	AddFunction(name string, f TypeDefinition) (string, error)
	NextScopeLayer() FunctionMap
}

func ToRpn(line string, fm FunctionMap) ([]string, parser.SyntaxError) {
	return toRpn(line, []rpnValue{}, []rpnValue{}, fm)
}

func ToGoExpression(line string, fm FunctionMap) (string, TypeDefinition, parser.SyntaxError) {
	rpn, err := ToRpn(line, fm)
	if err != nil {
		return "", nil, err
	}

	return ToInfix(rpn, fm)
}

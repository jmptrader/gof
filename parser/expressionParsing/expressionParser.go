package expressionParsing

import "github.com/apoydence/gof/parser"

type FunctionMap interface {
	GetFunction(name string) TypeDefinition
	AddFunction(name string, f TypeDefinition) (string, error)
	NextScopeLayer() FunctionMap
	NumberOfFunctions() int
}

func ToRpn(line string, lineNum int, fm FunctionMap) ([]RpnValue, parser.SyntaxError) {
	rpn, err := toRpn(line, []RpnValue{}, []RpnValue{}, fm)
	if err != nil {
		return nil, parser.NewSyntaxError(err.Error(), lineNum, 0)
	}

	return rpn, nil
}

func ToGoExpression(line string, lineNum int, fm FunctionMap) (string, TypeDefinition, parser.SyntaxError) {
	rpn, synErr := ToRpn(line, lineNum, fm)
	if synErr != nil {
		return "", nil, synErr
	}

	result, td, err := ToInfix(rpn, fm)
	if err != nil {
		return "", nil, parser.NewSyntaxError(err.Error(), lineNum, 0)
	}

	return result, td, nil
}

package expressionParsing

type FunctionMap interface {
	GetFunction(name string) *FuncTypeDefinition
	AddFunction(name string, f *FuncTypeDefinition) (string, error)
	NextScopeLayer() FunctionMap
}

func ToRpn(line string, fm FunctionMap) ([]string, error) {
	return toRpn(line, []rpnValue{}, []rpnValue{}, fm)
}

func ToGoExpression(line string, fm FunctionMap) (string, TypeDefinition, error) {
	rpn, err := ToRpn(line, fm)
	if err != nil {
		return "", nil, err
	}

	return ToInfix(rpn, fm)
}

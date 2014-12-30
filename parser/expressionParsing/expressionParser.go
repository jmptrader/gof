package expressionParsing

type FunctionMap interface {
	GetFunction(name string) *FunctionDeclaration
	AddFunction(name string, f *FunctionDeclaration) (string, error)
	NextScopeLayer() FunctionMap
}

func ToRpn(line string, fm FunctionMap) ([]string, error) {
	return toRpn(line, []rpnValue{}, []rpnValue{}, fm)
}

func ToGoExpression(line string, fm FunctionMap) (string, TypeName, error) {
	rpn, err := ToRpn(line, fm)
	if err != nil {
		return "", "", err
	}

	return ToInfix(rpn, fm)
}

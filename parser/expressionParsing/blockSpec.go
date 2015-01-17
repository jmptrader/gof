package expressionParsing

type blockSpec struct {
	block     string
	valueType TypeDefinition
}

func toBlockSpec(opQueue []string, fm FunctionMap) []*blockSpec {
	result := make([]*blockSpec, 0)
	for _, o := range opQueue {
		result = append(result, newBlockSpec(findType(o, fm)))
	}
	return result
}

func newBlockSpec(op string, valueType TypeDefinition) *blockSpec {
	return &blockSpec{
		block:     op,
		valueType: valueType,
	}
}

func findType(token string, fm FunctionMap) (string, TypeDefinition) {
	if fd := fm.GetFunction(token); fd != nil {
		return token, fd
	} else {
		return findPrimType(token)
	}
}

func findPrimType(token string) (string, TypeDefinition) {
	suffix1 := token[len(token)-1:]
	token1 := token[:len(token)-1]
	suffix2 := ""
	token2 := ""

	if len(token) > 2 {
		suffix2 = token[len(token)-2:]
		token2 = token[:len(token)-2]
	}

	if suffix2 == "ub" {
		return token2, safelyParseType("uint8")
	} else if suffix1 == "b" {
		return token1, safelyParseType("int8")
	} else if suffix2 == "uh" {
		return token2, safelyParseType("uint16")
	} else if suffix1 == "h" {
		return token1, safelyParseType("int16")
	} else if suffix2 == "ui" {
		return token2, safelyParseType("uint32")
	} else if suffix2 == "ul" {
		return token2, safelyParseType("uint64")
	} else if suffix1 == "l" {
		return token1, safelyParseType("int64")
	} else if suffix1 == "f" {
		return token1, safelyParseType("float32")
	} else if len(suffix2) > 0 && suffix2[0] == '.' {
		return token, safelyParseType("float64")
	} else {
		return token, safelyParseType("int32")
	}
}

func safelyParseType(s string) TypeDefinition {
	t, err, _ := ParseTypeDef(s)
	if err != nil {
		panic(err.Error())
	}
	return t
}

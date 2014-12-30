package expressionParsing

type blockSpec struct {
	block     string
	valueType TypeName
}

func toBlockSpec(opQueue []string, fm FunctionMap) []*blockSpec {
	result := make([]*blockSpec, 0)
	for _, o := range opQueue {
		result = append(result, newBlockSpec(findType(o, fm)))
	}
	return result
}

func newBlockSpec(op string, valueType TypeName) *blockSpec {
	return &blockSpec{
		block:     op,
		valueType: valueType,
	}
}

func findType(token string, fm FunctionMap) (string, TypeName) {
	if fd := fm.GetFunction(token); fd != nil {
		if fd.IsDefinition() {
			return fd.Name() + "()", fd.ReturnType()
		}
		return fd.Name(), fd.ReturnType()
	} else {
		return findPrimType(token)
	}
}

func findPrimType(token string) (string, TypeName) {
	suffix1 := token[len(token)-1:]
	token1 := token[:len(token)-1]
	suffix2 := ""
	token2 := ""

	if len(token) > 2 {
		suffix2 = token[len(token)-2:]
		token2 = token[:len(token)-2]
	}

	if suffix2 == "ub" {
		return token2, "uint8"
	} else if suffix1 == "b" {
		return token1, "int8"
	} else if suffix2 == "uh" {
		return token2, "uint16"
	} else if suffix1 == "h" {
		return token1, "int16"
	} else if suffix2 == "ui" {
		return token2, "uint32"
	} else if suffix2 == "ul" {
		return token2, "uint64"
	} else if suffix1 == "l" {
		return token1, "int64"
	} else if suffix1 == "f" {
		return token1, "float32"
	} else if len(suffix2) > 0 && suffix2[0] == '.' {
		return token, "float64"
	} else {
		return token, "int32"
	}
}

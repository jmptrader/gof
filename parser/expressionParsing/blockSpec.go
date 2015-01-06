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
		if tfd, ok := fd.(FuncTypeDefinition); ok && tfd.IsFunc() {
			if tfd.IsDefinition() {
				return tfd.FuncName() + "()", tfd.ReturnType()
			}
			return token, tfd.ReturnType()
		} else {
			return token, fd.ReturnType()
		}
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
		return token2, NewPrimTypeDefinition("uint8")
	} else if suffix1 == "b" {
		return token1, NewPrimTypeDefinition("int8")
	} else if suffix2 == "uh" {
		return token2, NewPrimTypeDefinition("uint16")
	} else if suffix1 == "h" {
		return token1, NewPrimTypeDefinition("int16")
	} else if suffix2 == "ui" {
		return token2, NewPrimTypeDefinition("uint32")
	} else if suffix2 == "ul" {
		return token2, NewPrimTypeDefinition("uint64")
	} else if suffix1 == "l" {
		return token1, NewPrimTypeDefinition("int64")
	} else if suffix1 == "f" {
		return token1, NewPrimTypeDefinition("float32")
	} else if len(suffix2) > 0 && suffix2[0] == '.' {
		return token, NewPrimTypeDefinition("float64")
	} else {
		return token, NewPrimTypeDefinition("int32")
	}
}

package expressionParsing

type blockSpec struct {
	block     string
	valueType TypeDefinition
	isArg     bool
}

func toBlockSpec(opQueue []RpnValue, fm FunctionMap) []*blockSpec {
	result := make([]*blockSpec, 0)
	for _, o := range opQueue {
		bs := newBlockSpec(findType(o, fm))
		bs.isArg = o.Argument
		result = append(result, bs)
	}

	establishFuncArgs(result, fm)
	return result
}

func newBlockSpec(op string, valueType TypeDefinition) *blockSpec {
	return &blockSpec{
		block:     op,
		valueType: valueType,
	}
}

func establishFuncArgs(result []*blockSpec, fm FunctionMap) {
	for i := len(result) - 1; i >= 1; i-- {
		for j := 0; i-j-1 >= 0 && j < argCount(result[i].valueType, 0); j++ {
			x := result[i-j-1]
			if _, ok := x.valueType.(FuncTypeDefinition); ok && x.isArg {
				x.valueType = NewArgTypeDefinition(x.valueType)
			}
		}
	}
}

func argCount(td TypeDefinition, count int) int {
	var ftd FuncTypeDefinition
	var ok bool
	if ftd, ok = td.(FuncTypeDefinition); !ok {
		return count
	}

	return argCount(ftd.ReturnType(), count+1)
}

func findType(rpn RpnValue, fm FunctionMap) (string, TypeDefinition) {
	if fd := fm.GetFunction(rpn.token); fd != nil {
		return rpn.token, fd
	} else {
		return findPrimType(rpn)
	}
}

func findPrimType(rpn RpnValue) (string, TypeDefinition) {
	suffix1 := rpn.token[len(rpn.token)-1:]
	token1 := rpn.token[:len(rpn.token)-1]
	suffix2 := ""
	token2 := ""

	if len(rpn.token) > 2 {
		suffix2 = rpn.token[len(rpn.token)-2:]
		token2 = rpn.token[:len(rpn.token)-2]
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
		return rpn.token, safelyParseType("float64")
	} else {
		return rpn.token, safelyParseType("int32")
	}
}

func safelyParseType(s string) TypeDefinition {
	t, err, _ := ParseTypeDef(s)
	if err != nil {
		panic(err.Error())
	}
	return t
}

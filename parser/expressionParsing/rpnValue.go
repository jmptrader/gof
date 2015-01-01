package expressionParsing

const (
	AddSub int = iota
	MultDiv
	FuncCall
)

type rpnValue struct {
	prec     int
	token    string
	leftPar  bool
	operator bool
}

func newPrimRpnValue(token string) rpnValue {
	return rpnValue{
		token: token,
	}
}

func newParenRpnValue() rpnValue {
	return rpnValue{
		leftPar: true,
		token:   ")",
	}
}

func newOpRpnValue(token string, prec int) rpnValue {
	return rpnValue{
		token:    token,
		operator: true,
		prec:     prec,
	}
}

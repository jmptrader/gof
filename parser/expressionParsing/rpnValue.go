package expressionParsing

const (
	AddSub int = iota
	MultDiv
	FuncCall
)

type RpnValue struct {
	prec     int
	token    string
	leftPar  bool
	Operator bool
	Argument bool
}

func NewPrimRpnValue(token string) RpnValue {
	return RpnValue{
		token:    token,
		Argument: true,
	}
}

func newParenRpnValue() RpnValue {
	return RpnValue{
		leftPar: true,
		token:   ")",
	}
}

func NewOpRpnValue(token string, prec int) RpnValue {
	return RpnValue{
		token:    token,
		Operator: true,
		prec:     prec,
	}
}

func (r RpnValue) String() string {
	return r.token
}

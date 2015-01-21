package expressionParsing

import (
	"github.com/apoydence/gof/parser"
	"strings"
)

const (
	AddSub int = iota
	MultDiv
	FuncCall
)

type RpnValue struct {
	prec     int
	token    string
	leftPar  bool
	operator bool
	argument bool
}

func newPrimRpnValue(token string) RpnValue {
	return RpnValue{
		token:    token,
		argument: true,
	}
}

func newParenRpnValue() RpnValue {
	return RpnValue{
		leftPar: true,
		token:   ")",
	}
}

func newOpRpnValue(token string, prec int) RpnValue {
	return RpnValue{
		token:    token,
		operator: true,
		prec:     prec,
	}
}

func (r RpnValue) String() string {
	return r.token
}

func ParseRpnValue(token string) RpnValue {
	isArg := strings.HasPrefix(token, "a:")
	if isArg {
		token = token[2:]
	}

	if parser.IsNumber(token) || (parser.ValidFunctionName(token) && isArg) {
		rpn := newPrimRpnValue(token)
		rpn.argument = isArg || parser.IsNumber(token)
		return rpn
	} else {
		var prec int
		if parser.ValidFunctionName(token) {
			prec = FuncCall
		} else {
			prec = opPrec(token)
		}
		return newOpRpnValue(token, prec)
	}
}

func ToRpnValues(tokens []string) []RpnValue {
	results := make([]RpnValue, 0)
	for _, token := range tokens {
		results = append(results, ParseRpnValue(token))
	}
	return results
}

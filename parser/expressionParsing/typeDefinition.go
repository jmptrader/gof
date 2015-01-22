package expressionParsing

import (
	"errors"
	"fmt"
	"github.com/apoydence/gof/parser"
)

type TypeDefinition interface {
	GenerateGo() string
}

func TypeDefEquals(a, b TypeDefinition) bool {
	return a.GenerateGo() == b.GenerateGo()
}

func ParseTypeDef(code string) (TypeDefinition, error, string) {
	first, rest := parser.Tokenize(code)
	_, rest2 := parser.Tokenize(rest)
	third, _ := parser.Tokenize(rest2)
	if first == "func" {
		argName, rest := parser.Tokenize(rest)

		if !parser.ValidFunctionName(argName) {
			return nil, errors.New(fmt.Sprintf("%s is not a valid argument name.", argName)), ""
		}

		if len(rest) == 0 {
			return nil, errors.New("Invalid function type definition"), ""
		}
		td, err, rest := ParseTypeDef(rest)
		if err != nil {
			return nil, err, ""
		}
		ret, err, rest := getRetType(rest)
		return newFuncTypeDefinition(argName, td, ret), nil, rest
	} else if third == "->" {
		return ParseTypeDef("func " + code)
	}

	return newPrimTypeDefinition(first), nil, rest
}

func getRetType(code string) (TypeDefinition, error, string) {
	first, rest := parser.Tokenize(code)
	if first == "->" {
		return ParseTypeDef(rest)
	}

	return nil, errors.New("Invalid function syntax"), ""
}

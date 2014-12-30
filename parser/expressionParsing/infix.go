package expressionParsing

import (
	"errors"
	"fmt"
	"github.com/apoydence/GoF/parser"
)

func ToInfix(opQueue []string, fm FunctionMap) (string, TypeName, error) {
	return toInfix(toBlockSpec(opQueue), fm, 0)
}

func toInfix(opQueue []*blockSpec, fm FunctionMap, index int) (string, TypeName, error) {
	if len(opQueue) <= index {
		return opQueue[0].block, opQueue[0].valueType, nil
	} else if parser.IsOperator(opQueue[index].block) {
		left, _ := addTypeToNumber(opQueue[index-2].block)
		right, _ := addTypeToNumber(opQueue[index-1].block)
		combined := fmt.Sprintf("(%s%s%s)", left, opQueue[index].block, right)
		value, err := getValueType(opQueue[index-2 : 3])
		if err != nil {
			return "", "", err
		}
		op := newBlockSpec(combined, value)
		return toInfix(append(opQueue[:index-2], combine(op, opQueue[index+1:])...), fm, index-2)
	} else {
		return toInfix(opQueue, fm, index+1)
	}
}

func getValueType(ops []*blockSpec) (TypeName, error) {
	left := ops[0].valueType
	right := ops[1].valueType
	if ops[0].valueType == "" {
		_, left = addTypeToNumber(ops[0].block)
	}
	if ops[1].valueType == "" {
		_, right = addTypeToNumber(ops[1].block)
	}
	if left != right {
		return "", errors.New(fmt.Sprintf("Illegal to %s%s%s", left, right, ops[2]))
	}

	return left, nil
}

func addTypeToNumber(token string) (string, TypeName) {
	if !parser.IsNumber(token) {
		return token, ""
	} else if len(token) <= 1 {
		return token, "int32"
	}

	suffix1 := token[len(token)-1:]
	token1 := token[:len(token)-1]
	suffix2 := ""
	token2 := ""

	if len(token) > 2 {
		suffix2 = token[len(token)-2:]
		token2 = token[:len(token)-2]
	}

	if suffix2 == "ub" {
		return fmt.Sprintf("uint8(%s)", token2), "uint8"
	} else if suffix1 == "b" {
		return fmt.Sprintf("int8(%s)", token1), "int8"
	} else if suffix2 == "uh" {
		return fmt.Sprintf("uint16(%s)", token2), "uint16"
	} else if suffix1 == "h" {
		return fmt.Sprintf("int16(%s)", token1), "int16"
	} else if suffix2 == "ui" {
		return fmt.Sprintf("uint32(%s)", token2), "uint32"
	} else if suffix2 == "ul" {
		return fmt.Sprintf("uint64(%s)", token2), "uint64"
	} else if suffix1 == "l" {
		return fmt.Sprintf("int64(%s)", token1), "int64"
	} else if suffix1 == "f" {
		return fmt.Sprintf("float32(%s)", token1), "float32"
	} else if len(suffix2) > 0 && suffix2[0] == '.' {
		return fmt.Sprintf("float64(%s)", token), "float64"
	} else {
		return token, "int32"
	}
}

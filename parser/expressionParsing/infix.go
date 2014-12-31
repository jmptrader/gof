package expressionParsing

import (
	"errors"
	"fmt"
	"github.com/apoydence/GoF/parser"
)

func ToInfix(opQueue []string, fm FunctionMap) (string, TypeDefinition, error) {
	return toInfix(toBlockSpec(opQueue, fm), fm, 0)
}

func toInfix(opQueue []*blockSpec, fm FunctionMap, index int) (string, TypeDefinition, error) {
	if len(opQueue) <= index {
		return opQueue[0].block, opQueue[0].valueType, nil
	} else if fd := fm.GetFunction(opQueue[index].block); fd != nil {
		arg := addTypeToNumber(opQueue[index-1])
		f := opQueue[index]
		combined := combine(newBlockSpec(fmt.Sprintf("%s(%s)", fd.FuncName(), arg), f.valueType), opQueue[index+1:])
		return toInfix(append(opQueue[:index-1], combined...), fm, index)
	} else if parser.IsOperator(opQueue[index].block) {
		left := addTypeToNumber(opQueue[index-2])
		right := addTypeToNumber(opQueue[index-1])
		combined := fmt.Sprintf("(%s%s%s)", left, opQueue[index].block, right)
		value, err := getValueType(opQueue[index-2 : 3])
		if err != nil {
			return "", nil, err
		}
		op := newBlockSpec(combined, value)
		return toInfix(append(opQueue[:index-2], combine(op, opQueue[index+1:])...), fm, index-2)
	} else {
		return toInfix(opQueue, fm, index+1)
	}
}

func getValueType(ops []*blockSpec) (TypeDefinition, error) {
	left := ops[0].valueType
	right := ops[1].valueType

	if left != right {
		return nil, errors.New(fmt.Sprintf("Illegal to %s%s%s", left, right, ops[2]))
	}

	return left, nil
}

func addTypeToNumber(block *blockSpec) string {
	if !parser.IsNumber(block.block) {
		return block.block
	}

	switch block.valueType.Name() {
	case "uint8":
		return fmt.Sprintf("uint8(%s)", block.block)
	case "int8":
		return fmt.Sprintf("int8(%s)", block.block)
	case "uint16":
		return fmt.Sprintf("uint16(%s)", block.block)
	case "int16":
		return fmt.Sprintf("int16(%s)", block.block)
	case "uint32":
		return fmt.Sprintf("uint32(%s)", block.block)
	case "int32":
		return block.block
	case "uint64":
		return fmt.Sprintf("uint64(%s)", block.block)
	case "int64":
		return fmt.Sprintf("int64(%s)", block.block)
	case "float32":
		return fmt.Sprintf("float32(%s)", block.block)
	case "float64":
		return fmt.Sprintf("float64(%s)", block.block)
	default:
		return block.block
	}
}

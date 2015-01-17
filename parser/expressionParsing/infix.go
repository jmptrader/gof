package expressionParsing

import (
	"errors"
	"fmt"
	"github.com/apoydence/gof/parser"
)

func ToInfix(opQueue []string, fm FunctionMap) (string, TypeDefinition, error) {
	return toInfix(toBlockSpec(opQueue, fm), fm, 0)
}

func toInfix(opQueue []*blockSpec, fm FunctionMap, index int) (string, TypeDefinition, error) {
	if len(opQueue) <= index {
		return opQueue[0].block, opQueue[0].valueType, nil
	} else if ftd, ok := isFunc(opQueue[index], fm); ok {
		argIndex := extractArgIndex(ftd, opQueue, index)
		arg := addTypeToNumber(opQueue[argIndex])
		ns := newBlockSpec(fmt.Sprintf("%s(%s)", opQueue[index].block, arg), ftd.ReturnType())
		combined := combine2(opQueue[argIndex+1:index], ns, opQueue[index+1:])
		return toInfix(append(opQueue[:argIndex], combined...), fm, argIndex)
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
	} else if td := fm.GetFunction(opQueue[index].block); td != nil {
		if _, ok := td.(FuncTypeDefinition); ok {
			panic("This block should NOT be a function. A previous if should have grabbed it")
		}
		opQueue[index] = newBlockSpec(fmt.Sprintf("%s()", opQueue[index].block), td)
		return toInfix(opQueue, fm, index+1)
	} else {
		return toInfix(opQueue, fm, index+1)
	}
}

func getFunction(block string, fm FunctionMap) (FuncTypeDefinition, bool) {
	if fd := fm.GetFunction(block); fd != nil {
		if ftd, ok := fd.(FuncTypeDefinition); ok {
			return ftd, true
		}
	}

	return FuncTypeDefinition{}, false
}

func isFunc(block *blockSpec, fm FunctionMap) (FuncTypeDefinition, bool) {
	if ftd, ok := block.valueType.(FuncTypeDefinition); ok {
		return ftd, true
	} else if ftd, ok := getFunction(block.block, fm); ok {
		return ftd, true
	}

	return FuncTypeDefinition{}, false
}

func extractArgIndex(ft FuncTypeDefinition, opQueue []*blockSpec, index int) int {
	count := 1
	fd := ft
	for {
		if _, ok := fd.ReturnType().(FuncTypeDefinition); !ok {
			break
		}
		count++
		fd = fd.ReturnType().(FuncTypeDefinition)
	}
	return index - count
}

func getValueType(ops []*blockSpec) (TypeDefinition, error) {
	left := ops[0]
	right := ops[1]

	if left.valueType != right.valueType {
		return nil, errors.New(fmt.Sprintf("Illegal to %s%s%s", left.block, ops[2].block, right.block))
	}

	return left.valueType, nil
}

func addTypeToNumber(block *blockSpec) string {
	if !parser.IsNumber(block.block) {
		return block.block
	}

	switch block.valueType.GenerateGo() {
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

func combine(x *blockSpec, y []*blockSpec) []*blockSpec {
	return append([]*blockSpec{x}, y...)
}

func combine2(otherArgs []*blockSpec, newBlock *blockSpec, rest []*blockSpec) []*blockSpec {
	return append(append(otherArgs, newBlock), rest...)
}

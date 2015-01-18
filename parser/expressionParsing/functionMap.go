package expressionParsing

import (
	"errors"
	"strconv"
)

type funcMap struct {
	prevScope FunctionMap
	fm        map[string]TypeDefinition
	count     int
}

func NewFunctionMap() FunctionMap {
	return &funcMap{
		fm: make(map[string]TypeDefinition),
	}
}

func newFunctionMap(prevScope FunctionMap) FunctionMap {
	return &funcMap{
		fm:        make(map[string]TypeDefinition),
		prevScope: prevScope,
	}
}

func (fm *funcMap) GetFunction(name string) TypeDefinition {
	if d, ok := fm.fm[name]; ok {
		return d
	} else if fm.prevScope != nil {
		return fm.prevScope.GetFunction(name)
	}
	return nil
}

func (fm *funcMap) AddFunction(name string, f TypeDefinition) (string, error) {
	if fm.GetFunction(name) != nil {
		return "", errors.New("The function name'" + name + "' is already allocated.")
	}

	if fd, ok := f.(FuncTypeDefinition); ok {
		f = fd
	}

	fm.fm[name] = f

	return name, nil
}
func (fm *funcMap) AdjustFunction(name string, f TypeDefinition) error {
	if _, ok := fm.fm[name]; !ok {
		return errors.New("Unknown function name: " + name)
	}

	fm.fm[name] = f

	return nil
}

func (fm *funcMap) NextScopeLayer() FunctionMap {
	return newFunctionMap(fm)
}

func (fm *funcMap) getNextFunctionName() string {
	goFuncName := "_" + strconv.Itoa(fm.count)
	fm.count++
	return goFuncName
}

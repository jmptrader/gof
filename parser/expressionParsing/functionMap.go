package expressionParsing

import (
	"errors"
	"strconv"
)

type funcMap struct {
	prevScope FunctionMap
	fm        map[string]*FuncTypeDefinition
	count     int
}

func NewFunctionMap() FunctionMap {
	return &funcMap{
		fm: make(map[string]*FuncTypeDefinition),
	}
}

func newFunctionMap(prevScope FunctionMap) FunctionMap {
	return &funcMap{
		fm:        make(map[string]*FuncTypeDefinition),
		prevScope: prevScope,
	}
}

func (fm *funcMap) GetFunction(name string) *FuncTypeDefinition {
	if d, ok := fm.fm[name]; ok {
		return d
	} else if fm.prevScope != nil {
		return fm.prevScope.GetFunction(name)
	}
	return nil
}

func (fm *funcMap) AddFunction(name string, f *FuncTypeDefinition) (string, error) {
	if fm.GetFunction(name) != nil {
		return "", errors.New("The function name'" + name + "' is already allocated.")
	}

	fm.fm[name] = f
	genName := fm.getNextFunctionName()
	f.name = genName

	return genName, nil
}

func (fm *funcMap) NextScopeLayer() FunctionMap {
	return newFunctionMap(fm)
}

func (fm *funcMap) getNextFunctionName() string {
	goFuncName := "_" + strconv.Itoa(fm.count)
	fm.count++
	return goFuncName
}

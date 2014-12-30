package expressionParsing

import (
	"errors"
	"strconv"
)

type funcMap struct {
	prevScope FunctionMap
	fm        map[string]*FunctionDeclaration
	count     int
}

func NewFunctionMap() FunctionMap {
	return &funcMap{
		fm: make(map[string]*FunctionDeclaration),
	}
}

func newFunctionMap(prevScope FunctionMap) FunctionMap {
	return &funcMap{
		fm:        make(map[string]*FunctionDeclaration),
		prevScope: prevScope,
	}
}

func (fm *funcMap) GetFunction(name string) *FunctionDeclaration {
	if d, ok := fm.fm[name]; ok {
		return d
	} else if fm.prevScope != nil {
		return fm.prevScope.GetFunction(name)
	}
	return nil
}

func (fm *funcMap) AddFunction(name string, f *FunctionDeclaration) (string, error) {
	if fm.GetFunction(name) != nil {
		return "", errors.New("The function name'" + name + "' is already allocated.")
	}

	fm.fm[name] = f
	return fm.getNextFunctionName(), nil
}

func (fm *funcMap) NextScopeLayer() FunctionMap {
	return newFunctionMap(fm)
}

func (fm *funcMap) getNextFunctionName() string {
	goFuncName := "_" + strconv.Itoa(fm.count)
	fm.count++
	return goFuncName
}

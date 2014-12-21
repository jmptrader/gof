package statementTypes

import (
	"errors"
	"strconv"
)

type funcMap struct {
	fm    map[string]FunctionDeclaration
	count int
}

func NewFunctionMap() FunctionMap {
	return &funcMap{
		fm: make(map[string]FunctionDeclaration),
	}
}

func (fm *funcMap) GetFunction(name string) FunctionDeclaration {
	if d, ok := fm.fm[name]; ok {
		return d
	}
	return nil
}

func (fm *funcMap) AddFunction(name string, f FunctionDeclaration) (string, error) {
	if _, ok := fm.fm[name]; ok {
		return "", errors.New("The function name'" + name + "' is already allocated.")
	}

	fm.fm[name] = f
	return fm.getNextFunctionName(), nil
}

func (fm *funcMap) getNextFunctionName() string {
	goFuncName := "_" + strconv.Itoa(fm.count)
	fm.count++
	return goFuncName
}

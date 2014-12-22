package statementTypes

import "github.com/apoydence/GoF/parser"

type TypeName string

type FunctionMap interface {
	GetFunction(name string) *FunctionDeclaration
	AddFunction(name string, f *FunctionDeclaration) (string, error)
}

type Statement interface {
	GenerateGo(fm FunctionMap) (string, TypeName, error)
}

type StatementParser interface {
	Parse(block string, nextBlockScanner *parser.ScanPeeker, factory *StatementFactory) Statement
}

type StatementFactory struct {
	statements []StatementParser
}

func NewStatementFactory(statements ...StatementParser) *StatementFactory {
	return &StatementFactory{
		statements: statements,
	}
}

func (sf *StatementFactory) Read(blockPeeker *parser.ScanPeeker) Statement {
	ok, value := blockPeeker.Read()

	if !ok {
		return nil
	}

	for _, s := range sf.statements {
		statement := s.Parse(value, blockPeeker, sf)
		if statement != nil {
			return statement
		}
	}

	return nil
}

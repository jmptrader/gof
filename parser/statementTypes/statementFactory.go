package statementTypes

import (
	"github.com/apoydence/GoF/parser"
	"github.com/apoydence/GoF/parser/expressionParsing"
)

type Statement interface {
	GenerateGo(fm expressionParsing.FunctionMap) (string, expressionParsing.TypeDefinition, parser.SyntaxError)
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

package statementTypes

import (
	"github.com/apoydence/gof/parser"
	"github.com/apoydence/gof/parser/expressionParsing"
)

type Statement interface {
	GenerateGo(fm expressionParsing.FunctionMap) (string, expressionParsing.TypeDefinition, parser.SyntaxError)
	LineNumber() int
}

type StatementParser interface {
	Parse(block string, lineNumber int, nextBlockScanner *parser.ScanPeeker, factory *StatementFactory) (Statement, parser.SyntaxError)
}

type StatementFactory struct {
	statements []StatementParser
}

func NewStatementFactory(statements ...StatementParser) *StatementFactory {
	return &StatementFactory{
		statements: statements,
	}
}

func (sf *StatementFactory) Read(blockPeeker *parser.ScanPeeker) (Statement, parser.SyntaxError) {
	ok, value, lineNum := blockPeeker.Read()

	if !ok {
		return nil, nil
	}

	for _, s := range sf.statements {
		statement, err := s.Parse(value, lineNum, blockPeeker, sf)
		if err != nil {
			return nil, err
		} else if statement != nil {
			return statement, nil
		}
	}

	return nil, nil
}

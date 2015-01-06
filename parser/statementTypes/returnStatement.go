package statementTypes

import (
	"github.com/apoydence/GoF/parser"
	"github.com/apoydence/GoF/parser/expressionParsing"
)

type ReturnStatement struct {
	block       string
	outputQueue []string
}

func NewReturnStatementParser() StatementParser {
	return ReturnStatement{}
}

func newReturnStatement(block string) Statement {
	statement := &ReturnStatement{
		block: block,
	}

	return statement
}

func (rs *ReturnStatement) OutputQueue() []string {
	return rs.outputQueue
}

func (rs ReturnStatement) Parse(block string, nextBlockScanner *parser.ScanPeeker, factory *StatementFactory) Statement {
	return newReturnStatement(block)
}

func (ds *ReturnStatement) GenerateGo(fm expressionParsing.FunctionMap) (string, expressionParsing.TypeDefinition, parser.SyntaxError) {
	return expressionParsing.ToGoExpression(ds.block, fm)
}

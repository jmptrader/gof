package statementTypes

import (
	"github.com/apoydence/gof/parser"
	"github.com/apoydence/gof/parser/expressionParsing"
)

type ReturnStatement struct {
	block       string
	outputQueue []string
	lineNum     int
}

func NewReturnStatementParser() StatementParser {
	return ReturnStatement{}
}

func newReturnStatement(block string, lineNum int) Statement {
	statement := &ReturnStatement{
		block:   block,
		lineNum: lineNum,
	}

	return statement
}

func (rs *ReturnStatement) OutputQueue() []string {
	return rs.outputQueue
}

func (rs ReturnStatement) Parse(block string, lineNum int, nextBlockScanner *parser.ScanPeeker, factory *StatementFactory) (Statement, parser.SyntaxError) {
	return newReturnStatement(block, lineNum), nil
}

func (ds *ReturnStatement) GenerateGo(fm expressionParsing.FunctionMap) (string, expressionParsing.TypeDefinition, parser.SyntaxError) {
	return expressionParsing.ToGoExpression(ds.block, fm)
}

func (ds *ReturnStatement) LineNumber() int {
	return ds.lineNum
}

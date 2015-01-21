package statementTypes

import (
	"fmt"
	"github.com/apoydence/gof/parser"
	"github.com/apoydence/gof/parser/expressionParsing"
	"regexp"
	"strings"
)

var innerStatementRegexp *regexp.Regexp = regexp.MustCompile("\\([\\w\\s+-/*>]+\\)")

type ReturnStatement struct {
	block       string
	outputQueue []string
	lineNum     int
	inner       []Statement
}

func NewReturnStatementParser() StatementParser {
	return ReturnStatement{}
}

func newReturnStatement(block string, lineNum int, inner []Statement) Statement {
	statement := &ReturnStatement{
		block:   block,
		lineNum: lineNum,
		inner:   inner,
	}

	return statement
}

func (rs *ReturnStatement) OutputQueue() []string {
	return rs.outputQueue
}

func (rs ReturnStatement) Parse(block string, lineNum int, nextBlockScanner *parser.ScanPeeker, factory *StatementFactory) (Statement, parser.SyntaxError) {
	b, inner, err := fetchInlineStatements(block, factory, lineNum)
	if err != nil {
		return nil, err
	}
	return newReturnStatement(b, lineNum, inner), nil
}

func (ds *ReturnStatement) GenerateGo(fm expressionParsing.FunctionMap) (string, expressionParsing.TypeDefinition, parser.SyntaxError) {

	mapper, err := generateGoFromInner(ds.inner, fm)
	if err != nil {
		return "", nil, err
	}
	code, td, err := expressionParsing.ToGoExpression(ds.block, ds.lineNum, fm)
	if err != nil {
		return "", nil, err
	}
	for k, v := range mapper {
		code = strings.Replace(code, k, v, 1)
	}
	return code, td, nil
}

func (ds *ReturnStatement) LineNumber() int {
	return ds.lineNum
}

func generateGoFromInner(statements []Statement, fm expressionParsing.FunctionMap) (map[string]string, parser.SyntaxError) {
	mapper := make(map[string]string)
	code := ""

	for i, s := range statements {
		name := fmt.Sprintf("x%d", i)
		c, td, err := s.GenerateGo(fm)
		if err != nil {
			return nil, err
		}
		fm.AddFunction(name, expressionParsing.NewArgTypeDefinition(td))
		mapper[name] = c
		code += fmt.Sprintf("%s\n", c)
	}

	return mapper, nil
}

func fetchInlineStatements(block string, factory *StatementFactory, lineNum int) (string, []Statement, parser.SyntaxError) {
	innerStats := make([]Statement, 0)
	count := 0
	for _, s := range innerStatementRegexp.FindAllString(block, -1) {
		code := s[1 : len(s)-1]
		peeker := parser.NewScanPeekerStr(code, lineNum)
		statement, err := factory.Read(peeker)
		if err != nil {
			return "", nil, err
		}
		if _, ok := statement.(*ReturnStatement); !ok {
			innerStats = append(innerStats, statement)
			key := fmt.Sprintf("( x%d )", count)
			count++
			block = strings.Replace(block, s, key, 1)
		}
	}
	return block, innerStats, nil
}

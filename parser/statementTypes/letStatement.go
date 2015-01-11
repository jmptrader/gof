package statementTypes

import (
	"fmt"
	"github.com/apoydence/gof/parser"
	"github.com/apoydence/gof/parser/expressionParsing"
)

type LetStatement struct {
	varName        string
	code           string
	innerStatement Statement
	lineNum        int
}

func NewLetParser() StatementParser {
	return LetStatement{}
}

func newLetStatement(varName, code string, lineNum int, innerStatement Statement) Statement {
	return &LetStatement{
		varName:        varName,
		code:           code,
		innerStatement: innerStatement,
		lineNum:        lineNum,
	}
}

func (ds LetStatement) Parse(block string, lineNum int, nextBlockScanner *parser.ScanPeeker, factory *StatementFactory) (Statement, parser.SyntaxError) {
	lines := parser.Lines(block)

	ok, varName, restOfLine := splitEquals(lines[0])

	if ok {
		//combinedLine := combineBlock(restOfLine, lines[1:])
		combinedLine := parser.FromLines(append([]string{restOfLine}, lines[1:]...))
		peeker := parser.NewScanPeekerStr(combinedLine, lineNum)
		st, err := factory.Read(peeker)
		if err != nil {
			return nil, err
		}
		return newLetStatement(varName, combinedLine, lineNum, st), nil
	}

	return nil, nil
}

func splitEquals(line string) (bool, string, string) {
	var varName string
	var rest string
	var let string
	var equals string
	let, rest = parser.Tokenize(line)
	varName, rest = parser.Tokenize(rest)
	equals, rest = parser.Tokenize(rest)

	if let == "let" && equals == "=" {
		return true, varName, rest
	}

	return false, "", ""
}

func combineBlock(firstLine string, lines []string) string {
	result := firstLine
	for _, line := range lines {
		if len(line) > 0 {
			if line[0] == '\t' {
				result = result + " " + line[1:]
			} else {
				break
			}
		}
	}
	return result
}

func (ds *LetStatement) GenerateGo(fm expressionParsing.FunctionMap) (string, expressionParsing.TypeDefinition, parser.SyntaxError) {
	innerCode, returnType, synErr := ds.innerStatement.GenerateGo(fm.NextScopeLayer())

	if synErr != nil {
		return "", nil, synErr
	}

	var fd expressionParsing.FuncTypeDefinition
	if returnType.IsFunc() {
		fd = returnType.(expressionParsing.FuncTypeDefinition)
	} else {
		fd = expressionParsing.NewFuncTypeDefinition("", nil, returnType)
	}

	name, err := fm.AddFunction(ds.varName, fd)

	if err != nil {
		return "", nil, parser.NewSyntaxError(err.Error(), 0, 0)
	}

	var genCode string
	if returnType.IsFunc() {
		genCode = fmt.Sprintf("var %s %s\n%s = %s", name, returnType.GenGo(), name, innerCode)
	} else {
		genCode = fmt.Sprintf("var %s func() %s\n%s = func(){\n\treturn %s\n}", name, returnType.GenGo(), name, innerCode)
	}
	return genCode, returnType, nil
}

func (ds *LetStatement) VariableName() string {
	return ds.varName
}

func (ds *LetStatement) CodeBlock() string {
	return ds.code
}

func (ds *LetStatement) LineNumber() int {
	return ds.lineNum
}

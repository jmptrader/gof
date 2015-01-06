package statementTypes

import (
	"fmt"
	"github.com/apoydence/GoF/parser"
	"github.com/apoydence/GoF/parser/expressionParsing"
)

type DeclarationStatement struct {
	varName        string
	code           string
	innerStatement Statement
}

func NewDeclarationParser() StatementParser {
	return DeclarationStatement{}
}

func newDeclarationStatement(varName, code string, innerStatement Statement) Statement {
	return &DeclarationStatement{
		varName:        varName,
		code:           code,
		innerStatement: innerStatement,
	}
}

func (ds DeclarationStatement) Parse(block string, nextBlockScanner *parser.ScanPeeker, factory *StatementFactory) Statement {
	lines := parser.Lines(block)

	ok, varName, restOfLine := splitEquals(lines[0])

	if ok {
		combinedLine := combineBlock(restOfLine, lines[1:])
		peeker := parser.NewScanPeekerStr(combinedLine)
		return newDeclarationStatement(varName, combinedLine, factory.Read(peeker))
	}

	return nil
}

func splitEquals(line string) (bool, string, string) {
	var varName string
	var rest string
	var equals string
	varName, rest = parser.Tokenize(line)
	equals, rest = parser.Tokenize(rest)

	if equals == "=" {
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

func (ds *DeclarationStatement) GenerateGo(fm expressionParsing.FunctionMap) (string, expressionParsing.TypeDefinition, parser.SyntaxError) {
	innerCode, returnType, synErr := ds.innerStatement.GenerateGo(fm.NextScopeLayer())

	if synErr != nil {
		return "", nil, synErr
	}

	fd := expressionParsing.NewFuncTypeDefinition("", nil, returnType)
	name, err := fm.AddFunction(ds.varName, fd)

	if err != nil {
		return "", nil, parser.NewSyntaxError(err.Error(), 0, 0)
	}

	genCode := fmt.Sprintf("var %s func() %s\n%s = func(){\n\treturn %s\n}", name, returnType.Name(), name, innerCode)
	return genCode, returnType, nil
}

func (ds *DeclarationStatement) VariableName() string {
	return ds.varName
}

func (ds *DeclarationStatement) CodeBlock() string {
	return ds.code
}

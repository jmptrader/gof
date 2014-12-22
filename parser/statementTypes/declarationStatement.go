package statementTypes

import (
	"github.com/apoydence/GoF/parser"
)

type DeclarationStatement struct {
	varName string
	block   string
}

func NewDeclarationParser() StatementParser {
	return DeclarationStatement{}
}

func newDeclarationStatement(varName, block string) Statement {
	return &DeclarationStatement{
		varName: varName,
		block:   block,
	}
}

func (ds DeclarationStatement) Parse(block string, nextBlockScanner *parser.ScanPeeker, factory *StatementFactory) Statement {
	lines := parser.Lines(block)

	ok, varName, restOfLine := splitEquals(lines[0])

	if ok {
		return newDeclarationStatement(varName, combineBlock(restOfLine, lines[1:]))
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

func (ds *DeclarationStatement) GenerateGo(fm FunctionMap) (string, TypeName, error) {
	return "", "", nil
}

func (ds *DeclarationStatement) VariableName() string {
	return ds.varName
}

func (ds *DeclarationStatement) CodeBlock() string {
	return ds.block
}

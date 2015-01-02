package statementTypes

import (
	"fmt"
	"github.com/apoydence/GoF/parser"
	"github.com/apoydence/GoF/parser/expressionParsing"
	"regexp"
)

var funcDeclRegex *regexp.Regexp

func init() {
	funcDeclRegex = regexp.MustCompile("^func\\s+(?P<name>[a-zA-Z]\\w*)\\s+(?P<typeDef>((\\s*->\\s*[a-zA-Z]\\w*\\s+[a-zA-Z]\\w*)+)((\\s+->\\s*[a-zA-Z]\\w*)))$")
}

type FunctionStatement struct {
	FuncName       string
	TypeDef        expressionParsing.TypeDefinition
	InnerStatement Statement
}

func NewFunctionStatementParser() StatementParser {
	return FunctionStatement{}
}

func newFunctionStatement(name string, typeDef *expressionParsing.FuncTypeDefinition, inner Statement) Statement {
	return &FunctionStatement{
		FuncName:       name,
		TypeDef:        typeDef,
		InnerStatement: inner,
	}
}

func (fs FunctionStatement) Parse(block string, nextBlockScanner *parser.ScanPeeker, factory *StatementFactory) Statement {
	lines := parser.Lines(block)
	name, typeDefStr, ok := fetchParts(lines[0])
	if !ok {
		return nil
	}

	typeDef, err := expressionParsing.ParseFuncTypeDefinition(typeDefStr)
	if err != nil {
		return nil
	}

	scanner := parser.NewScanPeekerStr(parser.FromLines(lines[1:]))
	innerStatement := factory.Read(scanner)

	return newFunctionStatement(name, typeDef, innerStatement)
}

func fetchParts(code string) (string, string, bool) {
	funcDeclRegex := regexp.MustCompile("^func\\s+(?P<name>[a-zA-Z]\\w*)\\s+(?P<typeDef>((\\s*->\\s*[a-zA-Z]\\w*\\s+[a-zA-Z]\\w*)+)((\\s+->\\s*[a-zA-Z]\\w*)))$")
	match := funcDeclRegex.FindStringSubmatch(code)
	groupIndex := make(map[string]int)
	for i, name := range funcDeclRegex.SubexpNames() {
		groupIndex[name] = i
	}

	if match == nil {
		return "", "", false
	}

	return match[groupIndex["name"]], match[groupIndex["typeDef"]], true
}

func (fs *FunctionStatement) GenerateGo(fm expressionParsing.FunctionMap) (string, expressionParsing.TypeDefinition, error) {
	return fmt.Sprintf(""), fs.TypeDef, nil
}

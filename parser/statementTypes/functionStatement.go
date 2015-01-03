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
	setupFuncMap(fm, fs.TypeDef.(*expressionParsing.FuncTypeDefinition))
	inner, _, _ := fs.InnerStatement.GenerateGo(fm.NextScopeLayer())
	return fmt.Sprintf("func %s %s{\n\t%s\n}", fs.FuncName, generateTypeDef(true, fs.TypeDef), generateInnerFunc(fs.TypeDef, 1, inner)), fs.TypeDef, nil
}

func setupFuncMap(fm expressionParsing.FunctionMap, typeDef *expressionParsing.FuncTypeDefinition) {
	if ft, ok := typeDef.ReturnType().(*expressionParsing.FuncTypeDefinition); ok {
		setupFuncMap(fm, ft)
	}
	newFt := expressionParsing.NewFuncTypeDefinition("", nil, typeDef.Argument)
	fm.AddFunction(typeDef.ArgumentName, newFt)
}

func generateInnerFunc(typeDef expressionParsing.TypeDefinition, tabCount int, innerStatement string) string {
	if !typeDef.ReturnType().IsFunc() {
		return fmt.Sprintf("return %s", innerStatement)
	}
	tabs := ""
	for i := 0; i <= tabCount; i++ {
		tabs += "\t"
	}
	tabs2 := string(tabs[:len(tabs)-1])
	return fmt.Sprintf("return %s {\n%s%s\n%s}", generateTypeDef(false, typeDef.ReturnType()), tabs, generateInnerFunc(typeDef.ReturnType(), tabCount+1, innerStatement), tabs2)
}

func generateTypeDef(first bool, typeDef expressionParsing.TypeDefinition) string {
	if ftd, ok := typeDef.(*expressionParsing.FuncTypeDefinition); ok {
		s := fmt.Sprintf("(%s %s) %s", ftd.ArgumentName, ftd.Argument.Name(), generateTypeDef(false, ftd.ReturnType()))
		if !first {
			return "func " + s
		}
		return s
	} else {
		return string(typeDef.Name())
	}
}

package statementTypes

import (
	"fmt"
	"github.com/apoydence/gof/parser"
	"github.com/apoydence/gof/parser/expressionParsing"
	"regexp"
)

var funcDeclRegex *regexp.Regexp

func init() {
	funcDeclRegex = regexp.MustCompile("^func\\s+(?P<name>[a-zA-Z]\\w*)\\s+(?P<typeDef>((\\s*->\\s*[a-zA-Z]\\w*\\s+[a-zA-Z]\\w*)+)((\\s+->\\s*[a-zA-Z]\\w*)))$")
}

type LambdaStatement struct {
	FuncName        string
	TypeDef         expressionParsing.TypeDefinition
	InnerStatements []Statement
	lineNum         int
}

func NewLambdaStatementParser() StatementParser {
	return LambdaStatement{}
}

func newLambdaStatement(name string, lineNum int, typeDef expressionParsing.FuncTypeDefinition, inner []Statement) Statement {
	return &LambdaStatement{
		FuncName:        name,
		TypeDef:         typeDef,
		InnerStatements: inner,
		lineNum:         lineNum,
	}
}

func (fs LambdaStatement) Parse(block string, lineNum int, nextBlockScanner *parser.ScanPeeker, factory *StatementFactory) (Statement, parser.SyntaxError) {
	lines := parser.Lines(block)
	name, typeDefStr, ok := fetchParts(lines[0])
	if !ok {
		return nil, nil
	}

	typeDef, err := expressionParsing.ParseFuncTypeDefinition(typeDefStr)
	if err != nil {
		return nil, err
	}

	innerStatements, err := fetchInnerStatements(parser.RemoveTabs(lines[1:]), factory, lineNum+1)
	if err != nil {
		return nil, err
	}

	err = verifyInnerStatements(innerStatements, lineNum)
	if err != nil {
		return nil, err
	}

	return newLambdaStatement(name, lineNum, typeDef, innerStatements), nil
}

func verifyInnerStatements(innerStatements []Statement, line int) parser.SyntaxError {
	numOfStatements := len(innerStatements)
	if numOfStatements == 0 {
		return parser.NewSyntaxError("No inner statement found", line, 0)
	} else if _, ok := innerStatements[numOfStatements-1].(*ReturnStatement); !ok {
		return parser.NewSyntaxError("Last statement in function is not a returnable statement", line, 0)
	} else {
		for _, s := range innerStatements[:numOfStatements-1] {
			if _, ok := s.(*LetStatement); !ok {
				return parser.NewSyntaxError("Only the last statement in function can be a returnable statement", line, 0)
			}
		}
	}

	return nil
}

func fetchInnerStatements(lines []string, factory *StatementFactory, lineNum int) ([]Statement, parser.SyntaxError) {
	scanner := parser.NewScanPeekerStr(parser.FromLines(lines), lineNum)
	statements := make([]Statement, 0)
	next := func() (Statement, parser.SyntaxError) {
		return factory.Read(scanner)
	}

	return subFetchInnerStatements(next, statements)
}

func subFetchInnerStatements(next func() (Statement, parser.SyntaxError), statements []Statement) ([]Statement, parser.SyntaxError) {

	s, err := next()
	if err != nil {
		return nil, err
	} else if s != nil {
		return subFetchInnerStatements(next, append(statements, s))
	}

	return statements, nil
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

func (fs *LambdaStatement) GenerateGo(fm expressionParsing.FunctionMap) (string, expressionParsing.TypeDefinition, parser.SyntaxError) {
	fm.AddFunction(fs.FuncName, fs.TypeDef)
	innerScope := fm.NextScopeLayer()
	setupFuncMap(innerScope, fs.TypeDef.(expressionParsing.FuncTypeDefinition))
	inner, err := generateInnerGo(innerScope, fs.InnerStatements)
	if err != nil {
		return "", nil, err
	}
	return fmt.Sprintf("func %s %s{\n\t%s\n}", fs.FuncName, generateTypeDef(true, fs.TypeDef), generateInnerFunc(fs.TypeDef, 1, inner)), fs.TypeDef, nil
}

func (fs *LambdaStatement) LineNumber() int {
	return fs.lineNum
}

func generateInnerGo(fm expressionParsing.FunctionMap, statements []Statement) ([]string, parser.SyntaxError) {
	code := make([]string, 0)
	for _, s := range statements {
		c, _, err := s.GenerateGo(fm)
		if err != nil {
			return nil, err
		}
		code = append(code, c)
	}

	return code, nil
}

func setupFuncMap(fm expressionParsing.FunctionMap, typeDef expressionParsing.FuncTypeDefinition) {
	if ft, ok := typeDef.ReturnType().(expressionParsing.FuncTypeDefinition); ok {
		setupFuncMap(fm, ft)
	}
	newFt := expressionParsing.NewPrimTypeDefinition(typeDef.Argument().Name())
	fm.AddFunction(typeDef.ArgumentName(), newFt)
}

func generateInnerFunc(typeDef expressionParsing.TypeDefinition, tabCount int, innerStatements []string) string {
	tabs := ""
	for i := 0; i <= tabCount; i++ {
		tabs += "\t"
	}
	tabs2 := string(tabs[:len(tabs)-1])

	if !typeDef.ReturnType().IsFunc() {
		lenInner := len(innerStatements) - 1
		innerCode := parser.FromLines(innerStatements[:lenInner])
		return fmt.Sprintf("%s\n%sreturn %s", innerCode, tabs2, innerStatements[lenInner])
	}

	return fmt.Sprintf("return %s {\n%s%s\n%s}", generateTypeDef(false, typeDef.ReturnType()), tabs, generateInnerFunc(typeDef.ReturnType(), tabCount+1, innerStatements), tabs2)
}

func generateTypeDef(first bool, typeDef expressionParsing.TypeDefinition) string {
	if ftd, ok := typeDef.(expressionParsing.FuncTypeDefinition); ok {
		s := fmt.Sprintf("(%s %s) %s", ftd.ArgumentName(), ftd.Argument().Name(), generateTypeDef(false, ftd.ReturnType()))
		if !first {
			return "func " + s
		}
		return s
	} else {
		return string(typeDef.Name())
	}
}

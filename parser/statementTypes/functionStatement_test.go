package statementTypes_test

import (
	"github.com/apoydence/GoF/parser/expressionParsing"
	. "github.com/apoydence/GoF/parser/statementTypes"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("FunctionStatement", func() {
	var statementParser StatementParser
	var factory *StatementFactory
	BeforeEach(func() {
		retParser := NewReturnStatementParser()
		declParser := NewDeclarationParser()
		statementParser = NewFunctionStatementParser()
		factory = NewStatementFactory(declParser, statementParser, retParser)
	})
	Context("Parse", func() {
		It("Should distinguish a function statement", func() {
			code := "func someName -> a A -> b B -> int32\n\t5+9"
			f, err := statementParser.Parse(code, nil, factory)
			fs := f.(*FunctionStatement)
			Expect(err).To(BeNil())
			Expect(fs).ToNot(BeNil())
			Expect(fs.FuncName).To(Equal("someName"))
			Expect(fs.TypeDef.Name()).To(BeEquivalentTo("a A->b B->int32"))
			Expect(fs.InnerStatements).ToNot(BeNil())
			Expect(fs.InnerStatements).To(HaveLen(1))
		})
	})
	Context("GenerateGo", func() {
		It("Should generate a proper Go function with one inner statement", func() {
			fm := expressionParsing.NewFunctionMap()
			code := "func addTogether -> a int32 -> b int32 -> c int32 -> int32\n\ta + b + c"
			f, err := statementParser.Parse(code, nil, factory)
			code, _, err = f.GenerateGo(fm)
			Expect(err).To(BeNil())
			code, _, err = f.GenerateGo(fm)
			Expect(err).To(BeNil())
			Expect(fm.GetFunction("addTogether")).ToNot(BeNil())
			Expect(matchCode(code, "func addTogether (a int32) func (b int32) func (c int32) int32{\n\treturn func (b int32) func (c int32) int32 {\n\t\treturn func (c int32) int32 {\n\t\t\t\n\t\t\treturn ((a+b)+c)\n\t\t}\n\t}\n}")).To(BeTrue())
		})
		It("Should generate a proper Go function with multiple inner statements", func() {
			fm := expressionParsing.NewFunctionMap()
			code := "func addTogether -> a int32 -> b int32 -> c int32 -> int32\n\td = 6\n\ta + b + c + d"
			f, err := statementParser.Parse(code, nil, factory)
			Expect(err).To(BeNil())
			code, _, err = f.GenerateGo(fm)
			Expect(err).To(BeNil())
			Expect(matchCode(code, "func addTogether (a int32) func (b int32) func (c int32) int32{\n\treturn func (b int32) func (c int32) int32 {\n\t\treturn func (c int32) int32 {\n\t\t\t\n\t\t\tvar d func() int32\nd = func(){return 6}return (((a+b)+c)+d())\n\t\t}\n\t}\n}")).To(BeTrue())
		})
	})
})

func matchCode(codeA, codeB string) bool {
	return stripWhitespace(codeA) == stripWhitespace(codeB)
}

func stripWhitespace(s string) string {
	noTabs := strings.Replace(s, "\t", "", -1)
	return strings.Replace(noTabs, "\n", "", -1)
}

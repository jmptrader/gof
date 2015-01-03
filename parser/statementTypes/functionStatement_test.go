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
			f := statementParser.Parse(code, nil, factory).(*FunctionStatement)
			Expect(f).ToNot(BeNil())
			Expect(f.FuncName).To(Equal("someName"))
			Expect(f.TypeDef.Name()).To(BeEquivalentTo("a A->b B->int32"))
			Expect(f.InnerStatements).ToNot(BeNil())
			Expect(f.InnerStatements).To(HaveLen(1))
		})
	})
	Context("GenerateGo", func() {
		It("Should generate a proper Go function with one inner statement", func() {
			fm := expressionParsing.NewFunctionMap()
			code := "func addTogether -> a int32 -> b int32 -> c int32 -> int32\n\ta + b + c"
			f := statementParser.Parse(code, nil, factory)
			code, _, err := f.GenerateGo(fm)
			Expect(err).To(BeNil())
			Expect(matchCode(code, "func addTogether (a int32) func (b int32) func (c int32) int32{\n\treturn func (b int32) func (c int32) int32 {\n\t\treturn func (c int32) int32 {\n\t\t\t\n\t\t\treturn ((_2()+_1())+_0())\n\t\t}\n\t}\n}")).To(BeTrue())
		})
		It("Should generate a proper Go function with multiple inner statements", func() {
			fm := expressionParsing.NewFunctionMap()
			code := "func addTogether -> a int32 -> b int32 -> c int32 -> int32\n\td = 6\n\ta + b + c + d"
			f := statementParser.Parse(code, nil, factory)
			code, _, err := f.GenerateGo(fm)
			Expect(err).To(BeNil())
			Expect(matchCode(code, "func addTogether (a int32) func (b int32) func (c int32) int32{\n\treturn func (b int32) func (c int32) int32 {\n\t\treturn func (c int32) int32 {\n\t\t\t\n\t\t\tvar _3 func() int32\n_3 = func(){return 6}return (((_2()+_1())+_0())+_3())\n\t\t}\n\t}\n}")).To(BeTrue())
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

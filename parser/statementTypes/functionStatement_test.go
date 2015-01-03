package statementTypes_test

import (
	"github.com/apoydence/GoF/parser/expressionParsing"
	. "github.com/apoydence/GoF/parser/statementTypes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("FunctionStatement", func() {
	var statementParser StatementParser
	var factory *StatementFactory
	BeforeEach(func() {
		retParser := NewReturnStatementParser()
		statementParser = NewFunctionStatementParser()
		factory = NewStatementFactory(statementParser, retParser)
	})
	Context("Parse", func() {
		It("Should distinguish a function statement", func() {
			code := "func someName -> a A -> b B -> int32\n\t5+9"
			f := statementParser.Parse(code, nil, factory).(*FunctionStatement)
			Expect(f).ToNot(BeNil())
			Expect(f.FuncName).To(Equal("someName"))
			Expect(f.TypeDef.Name()).To(BeEquivalentTo("a A->b B->int32"))
			Expect(f.InnerStatement).ToNot(BeNil())
		})
	})
	Context("GenerateGo", func() {
		It("Should generate a proper Go function", func() {
			fm := expressionParsing.NewFunctionMap()
			code := "func addTogether -> a int32 -> b int32 -> c int32 -> int32\n\ta + b + c"
			f := statementParser.Parse(code, nil, factory)
			code, _, err := f.GenerateGo(fm)
			Expect(err).To(BeNil())
			Expect(code).To(Equal("func addTogether (a int32) func (b int32) func (c int32) int32{\n\treturn func (b int32) func (c int32) int32 {\n\t\treturn func (c int32) int32 {\n\t\t\treturn ((_2()+_1())+_0())\n\t\t}\n\t}\n}"))
		})
	})
})

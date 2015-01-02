package statementTypes_test

import (
	. "github.com/apoydence/GoF/parser/statementTypes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("FunctionStatement", func() {
	Context("Parse", func() {
		var statementParser StatementParser
		var factory *StatementFactory
		BeforeEach(func() {
			retParser := NewReturnStatementParser()
			statementParser = NewFunctionStatementParser()
			factory = NewStatementFactory(statementParser, retParser)
		})
		It("Should distinguish a function statement", func() {
			code := "func someName -> a A -> b B -> int32\n\t5+9"
			f := statementParser.Parse(code, nil, factory).(*FunctionStatement)
			Expect(f).ToNot(BeNil())
			Expect(f.FuncName).To(Equal("someName"))
			Expect(f.TypeDef.Name()).To(BeEquivalentTo("A->B->int32"))
			Expect(f.InnerStatement).ToNot(BeNil())
		})
	})
	Context("GenerateGo", func() {
	})
})

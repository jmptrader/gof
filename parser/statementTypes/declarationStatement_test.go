package statementTypes_test

import (
	. "github.com/apoydence/GoF/parser/statementTypes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DeclarationStatement", func() {
	Context("Parse", func() {
		var statementParser StatementParser
		var factory *StatementFactory

		BeforeEach(func() {
			statementParser = NewDeclarationParser()
			factory = NewStatementFactory(statementParser)
		})

		It("Should pick out any declaration", func() {
			code := "a = b + 9"
			d := statementParser.Parse(code, nil, factory).(*DeclarationStatement)
			Expect(d).ToNot(BeNil())
			Expect(d.VariableName()).To(Equal("a"))
			Expect(d.CodeBlock()).To(Equal("b + 9"))
		})
		It("Should pick out a multi-lined declaration", func() {
			code := "a = b + 9\n\t+ 6"
			d := statementParser.Parse(code, nil, factory).(*DeclarationStatement)
			Expect(d).ToNot(BeNil())
			Expect(d.VariableName()).To(Equal("a"))
			Expect(d.CodeBlock()).To(Equal("b + 9 + 6"))
		})
		It("Should return nil for a non declaration", func() {
			code := "if true\n\t9\nelse\n\t10"
			d := statementParser.Parse(code, nil, factory)
			Expect(d).To(BeNil())
		})
	})
})

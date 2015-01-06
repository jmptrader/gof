package statementTypes_test

import (
	"github.com/apoydence/GoF/parser/expressionParsing"
	. "github.com/apoydence/GoF/parser/statementTypes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DeclarationStatement", func() {
	var statementParser StatementParser
	var factory *StatementFactory

	BeforeEach(func() {
		statementParser = NewDeclarationParser()
		returnStatement := NewReturnStatementParser()
		factory = NewStatementFactory(statementParser, returnStatement)
	})
	Context("Parse", func() {

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
	Context("GenerateGo", func() {
		It("Should return proper go code", func() {
			code := "a = 9 + 6"
			d := statementParser.Parse(code, nil, factory).(*DeclarationStatement)
			fm := expressionParsing.NewFunctionMap()
			gocode, returnType, err := d.GenerateGo(fm)
			Expect(err).To(BeNil())
			Expect(returnType.Name()).To(BeEquivalentTo("int32"))
			Expect(gocode).To(Equal("var a func() int32\na = func(){\n\treturn (9+6)\n}"))
			fd := fm.GetFunction("a")
			Expect(fd).ToNot(BeNil())
			Expect(fd.ReturnType().Name()).To(BeEquivalentTo("int32"))
			Expect(fd.(expressionParsing.FuncTypeDefinition).IsDefinition()).To(BeTrue())
		})
	})
})

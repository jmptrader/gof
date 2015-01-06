package statementTypes_test

import (
	"github.com/apoydence/gof/parser/expressionParsing"
	. "github.com/apoydence/gof/parser/statementTypes"

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
			d, err := statementParser.Parse(code, nil, factory)
			ds := d.(*DeclarationStatement)
			Expect(err).To(BeNil())
			Expect(ds).ToNot(BeNil())
			Expect(ds.VariableName()).To(Equal("a"))
			Expect(ds.CodeBlock()).To(Equal("b + 9"))
		})
		It("Should pick out a multi-lined declaration", func() {
			code := "a = b + 9\n\t+ 6"
			d, err := statementParser.Parse(code, nil, factory)
			ds := d.(*DeclarationStatement)
			Expect(err).To(BeNil())
			Expect(ds).ToNot(BeNil())
			Expect(ds.VariableName()).To(Equal("a"))
			Expect(ds.CodeBlock()).To(Equal("b + 9 + 6"))
		})
		It("Should return nil for a non declaration", func() {
			code := "if true\n\t9\nelse\n\t10"
			d, err := statementParser.Parse(code, nil, factory)
			Expect(err).To(BeNil())
			Expect(d).To(BeNil())
		})
	})
	Context("GenerateGo", func() {
		It("Should return proper go code", func() {
			code := "a = 9 + 6"
			d, err := statementParser.Parse(code, nil, factory)
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

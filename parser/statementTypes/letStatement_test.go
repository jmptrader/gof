package statementTypes_test

import (
	"github.com/apoydence/gof/parser/expressionParsing"
	. "github.com/apoydence/gof/parser/statementTypes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("LetStatement", func() {
	var statementParser StatementParser
	var factory *StatementFactory

	BeforeEach(func() {
		statementParser = NewLetParser()
		returnStatement := NewReturnStatementParser()
		factory = NewStatementFactory(statementParser, returnStatement)
	})
	Context("Parse", func() {

		It("Should pick out any declaration", func() {
			code := "let a = b + 9"
			d, err := statementParser.Parse(code, 0, nil, factory)
			ds := d.(*LetStatement)
			Expect(err).To(BeNil())
			Expect(ds).ToNot(BeNil())
			Expect(ds.VariableName()).To(Equal("a"))
			Expect(stripWhitespace(ds.CodeBlock())).To(Equal("b+9"))
		})
		It("Should return nil for a non declaration", func() {
			code := "if true\n\t9\nelse\n\t10"
			d, err := statementParser.Parse(code, 0, nil, factory)
			Expect(err).To(BeNil())
			Expect(d).To(BeNil())
		})
	})
	Context("GenerateGo", func() {
		It("Should return proper go code", func() {
			code := "let a = 9 + 6"
			d, err := statementParser.Parse(code, 0, nil, factory)
			fm := expressionParsing.NewFunctionMap()
			gocode, returnType, err := d.GenerateGo(fm)
			Expect(err).To(BeNil())
			intType, _, _ := expressionParsing.ParseTypeDef("int32")
			Expect(returnType.GenerateGo()).To(Equal(intType.GenerateGo()))
			Expect(gocode).To(Equal("var a func() int32\na = func(){\n\treturn (9+6)\n}"))
			fd := fm.GetFunction("a")
			Expect(fd).ToNot(BeNil())
			Expect(fd.GenerateGo()).To(Equal(intType.GenerateGo()))
		})
	})
})

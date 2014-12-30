package statementTypes_test

import (
	"fmt"
	"github.com/apoydence/GoF/parser/expressionParsing"
	. "github.com/apoydence/GoF/parser/statementTypes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ReturnStatement", func() {
	var statementParser StatementParser
	var factory *StatementFactory

	BeforeEach(func() {
		statementParser = NewReturnStatementParser()
		factory = NewStatementFactory(statementParser)
	})
	Context("Parse", func() {

		It("Should pick out the statement", func() {
			code := "a + 9"
			r := statementParser.Parse(code, nil, factory)
			Expect(r).ToNot(BeNil())
		})
	})
	Context("GenerateGo", func() {
		var funcMap expressionParsing.FunctionMap
		BeforeEach(func() {
			funcMap = expressionParsing.NewFunctionMap()
		})
		Context("No functions", func() {
			It("Should generate the proper Go code", func() {
				code := "( 1 + 11 ) - 10 / 2"
				r := statementParser.Parse(code, nil, factory)
				genGo, returnType, err := r.GenerateGo(funcMap)
				Expect(err).To(BeNil())
				Expect(returnType).To(BeEquivalentTo("int32"))
				Expect(genGo).To(Equal("((1+11)-(10/2))"))
			})
			It("Should generate the proper Go code with proper numeric types", func() {
				code := "( 1b + 11b ) - 10b / 2b"
				r := statementParser.Parse(code, nil, factory)
				genGo, returnType, err := r.GenerateGo(funcMap)
				Expect(err).To(BeNil())
				Expect(returnType).To(BeEquivalentTo("int8"))
				Expect(genGo).To(Equal("((int8(1)+int8(11))-(int8(10)/int8(2)))"))
			})
		})
		Context("With functions", func() {
			It("Should generate the proper Go code with a definition", func() {
				code := "( 1 + 11 ) - a / 2"
				r := statementParser.Parse(code, nil, factory)
				fm := expressionParsing.NewFunctionMap()
				name, _ := fm.AddFunction("a", expressionParsing.NewDefinition("int32"))
				genGo, returnType, err := r.GenerateGo(fm)
				Expect(err).To(BeNil())
				Expect(returnType).To(BeEquivalentTo("int32"))
				Expect(genGo).To(Equal(fmt.Sprintf("((1+11)-(%s()/2))", name)))
			})
			It("Should generate the proper Go code with a function", func() {
			})
		})
	})
})

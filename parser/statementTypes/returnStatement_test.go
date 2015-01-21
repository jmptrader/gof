package statementTypes_test

import (
	"fmt"
	"github.com/apoydence/gof/parser/expressionParsing"
	. "github.com/apoydence/gof/parser/statementTypes"

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
			r, err := statementParser.Parse(code, 0, nil, factory)
			Expect(err).To(BeNil())
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
				r, err := statementParser.Parse(code, 0, nil, factory)
				Expect(err).To(BeNil())
				genGo, returnType, err := r.GenerateGo(funcMap)
				Expect(err).To(BeNil())
				intType, _, _ := expressionParsing.ParseTypeDef("int32")
				Expect(returnType.GenerateGo()).To(Equal(intType.GenerateGo()))
				Expect(genGo).To(Equal("((1+11)-(10/2))"))
			})
			It("Should generate the proper Go code with proper numeric types", func() {
				code := "( 1b + 11b ) - 10b / 2b"
				r, err := statementParser.Parse(code, 0, nil, factory)
				Expect(err).To(BeNil())
				genGo, returnType, err := r.GenerateGo(funcMap)
				Expect(err).To(BeNil())
				intType, _, _ := expressionParsing.ParseTypeDef("int8")
				Expect(returnType.GenerateGo()).To(Equal(intType.GenerateGo()))
				Expect(genGo).To(Equal("((int8(1)+int8(11))-(int8(10)/int8(2)))"))
			})
		})
		Context("With functions", func() {
			var intType expressionParsing.TypeDefinition
			var funcType expressionParsing.TypeDefinition
			var twoArgFunc expressionParsing.TypeDefinition
			BeforeEach(func() {
				intType, _, _ = expressionParsing.ParseTypeDef("int32")
				funcType, _, _ = expressionParsing.ParseTypeDef("func a int32 -> int32")
				twoArgFunc, _, _ = expressionParsing.ParseTypeDef("func a int32 -> b int32 -> int32")
			})
			It("Should generate the proper Go code with a definition", func() {
				code := "( 1 + 11 ) - a / 2"
				r, err := statementParser.Parse(code, 0, nil, factory)
				Expect(err).To(BeNil())
				fm := expressionParsing.NewFunctionMap()
				name, _ := fm.AddFunction("a", intType)
				genGo, returnType, err := r.GenerateGo(fm)
				Expect(err).To(BeNil())
				Expect(returnType.GenerateGo()).To(Equal(intType.GenerateGo()))
				Expect(genGo).To(Equal(fmt.Sprintf("((1+11)-(%s()/2))", name)))
			})
			It("Should generate the proper Go code with definitions", func() {
				code := "a + b + c"
				r, err := statementParser.Parse(code, 0, nil, factory)
				Expect(err).To(BeNil())
				fm := expressionParsing.NewFunctionMap()
				nameA, _ := fm.AddFunction("a", intType)
				nameB, _ := fm.AddFunction("b", intType)
				nameC, _ := fm.AddFunction("c", intType)
				genGo, returnType, err := r.GenerateGo(fm)
				Expect(err).To(BeNil())
				Expect(returnType.GenerateGo()).To(Equal(intType.GenerateGo()))
				Expect(genGo).To(Equal(fmt.Sprintf("((%s()+%s())+%s())", nameA, nameB, nameC)))
			})
			It("Should generate the proper Go code with a function", func() {
				code := "( 7 + 13 ) - a 5 / 8"
				r, err := statementParser.Parse(code, 0, nil, factory)
				Expect(err).To(BeNil())
				fm := expressionParsing.NewFunctionMap()
				name, _ := fm.AddFunction("a", funcType)
				genGo, returnType, err := r.GenerateGo(fm)
				Expect(err).To(BeNil())
				Expect(returnType.GenerateGo()).To(Equal(intType.GenerateGo()))
				Expect(genGo).To(Equal(fmt.Sprintf("((7+13)-(%s(5)/8))", name)))
			})
			It("Should generate the proper Go code with a curried function", func() {
				code := "( 7 + 13 ) - a 5 9 / 8"
				r, err := statementParser.Parse(code, 0, nil, factory)
				Expect(err).To(BeNil())
				fm := expressionParsing.NewFunctionMap()
				name, _ := fm.AddFunction("a", twoArgFunc)
				genGo, returnType, err := r.GenerateGo(fm)
				Expect(err).To(BeNil())
				Expect(returnType.GenerateGo()).To(Equal(intType.GenerateGo()))
				Expect(genGo).To(Equal(fmt.Sprintf("((7+13)-(%s(5)(9)/8))", name)))
			})
			It("Should generate the proper Go code with multiple functions", func() {
				code := "( 7 + 13 ) - a 5 ( b 9 ) / 8"
				r, err := statementParser.Parse(code, 0, nil, factory)
				Expect(err).To(BeNil())
				fm := expressionParsing.NewFunctionMap()
				nameA, _ := fm.AddFunction("a", twoArgFunc)
				nameB, _ := fm.AddFunction("b", funcType)
				genGo, returnType, err := r.GenerateGo(fm)
				Expect(err).To(BeNil())
				Expect(returnType.GenerateGo()).To(Equal(intType.GenerateGo()))
				Expect(genGo).To(Equal(fmt.Sprintf("((7+13)-(%s(5)(%s(9))/8))", nameA, nameB)))
			})
			It("Should generate the proper Go code with layered functions", func() {
				code := "( 7 + 13 ) - c ( a 5 b 9 / 8 ) 10"
				r, err := statementParser.Parse(code, 0, nil, factory)
				Expect(err).To(BeNil())
				fm := expressionParsing.NewFunctionMap()
				nameA, _ := fm.AddFunction("a", twoArgFunc)
				nameB, _ := fm.AddFunction("b", funcType)
				nameC, _ := fm.AddFunction("c", twoArgFunc)
				genGo, returnType, err := r.GenerateGo(fm)
				Expect(err).To(BeNil())
				Expect(returnType.GenerateGo()).To(Equal(intType.GenerateGo()))
				Expect(genGo).To(Equal(fmt.Sprintf("((7+13)-%s((%s(5)(%s(9))/8))(10))", nameC, nameA, nameB)))
			})
			It("Should return an error when adding a double to an int", func() {
				code := " 7 + 13.0"
				r, err := statementParser.Parse(code, 9, nil, factory)
				Expect(err).To(BeNil())
				fm := expressionParsing.NewFunctionMap()
				_, _, err = r.GenerateGo(fm)
				Expect(err).ToNot(BeNil())
				Expect(err.Line()).To(Equal(9))
			})
		})
	})
})

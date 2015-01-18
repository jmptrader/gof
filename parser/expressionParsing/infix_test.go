package expressionParsing_test

import (
	. "github.com/apoydence/gof/parser/expressionParsing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Infix", func() {
	Context("No functions", func() {
		It("It should convert RPN to infix", func() {
			rpn := []string{"5", "6", "+", "7", "/"}
			fm := NewFunctionMap()
			code, td, err := ToInfix(rpn, fm)
			Expect(err).To(BeNil())
			Expect(td.GenerateGo()).To(Equal("int32"))
			Expect(code).To(Equal("((5+6)/7)"))
		})
	})
	Context("With functions", func() {
		It("Should use a definition as a normal number", func() {
			rpn := []string{"5", "6", "+", "a", "/"}
			fm := NewFunctionMap()
			intType, _, _ := ParseTypeDef("int32")
			fm.AddFunction("a", intType)
			code, td, err := ToInfix(rpn, fm)
			Expect(err).To(BeNil())
			Expect(td.GenerateGo()).To(Equal("int32"))
			Expect(code).To(Equal("((5+6)/a())"))
		})
		It("Should use a function with multiple arguments", func() {
			rpn := []string{"5", "6", "+", "7", "8", "a", "/"}
			fm := NewFunctionMap()
			f, _, _ := ParseTypeDef("func x int32 -> y int32 -> int32")
			fm.AddFunction("a", f)
			code, td, err := ToInfix(rpn, fm)
			Expect(err).To(BeNil())
			Expect(td.GenerateGo()).To(Equal("int32"))
			Expect(code).To(Equal("((5+6)/a(7)(8))"))
		})
		It("Should use a argument as a normal value", func() {
			rpn := []string{"5", "6", "+", "a", "/"}
			fm := NewFunctionMap()
			intType, _, _ := ParseTypeDef("int32")
			fm.AddFunction("a", NewArgTypeDefinition(intType.(PrimTypeDefinition)))
			code, td, err := ToInfix(rpn, fm)
			Expect(err).To(BeNil())
			Expect(td.GenerateGo()).To(Equal("int32"))
			Expect(code).To(Equal("((5+6)/a)"))
		})
	})
})

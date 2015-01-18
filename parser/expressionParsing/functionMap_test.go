package expressionParsing_test

import (
	. "github.com/apoydence/gof/parser/expressionParsing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("FunctionMap", func() {
	Context("Multiple Functions", func() {
		var intType TypeDefinition
		var funcType TypeDefinition
		BeforeEach(func() {
			intType, _, _ = ParseTypeDef("int32")
			funcType, _, _ = ParseTypeDef("func x int32 -> int32")
		})
		It("Should return incrementing values", func() {
			fm := NewFunctionMap()
			a, err := fm.AddFunction("a", funcType)
			Expect(err).To(BeNil())
			b, err := fm.AddFunction("b", funcType)
			Expect(err).To(BeNil())
			c, err := fm.AddFunction("c", funcType)
			Expect(err).To(BeNil())

			Expect(a).To(Equal("a"))
			Expect(b).To(Equal("b"))
			Expect(c).To(Equal("c"))

			Expect(fm.GetFunction("a")).To(Equal(funcType))
			Expect(fm.GetFunction("b")).To(Equal(funcType))
			Expect(fm.GetFunction("c")).To(Equal(funcType))
		})
		It("Should return null for an unknown function", func() {
			fm := NewFunctionMap()
			Expect(fm.GetFunction("x")).To(BeNil())
		})
		It("Should return an error when add the same function name", func() {
			fm := NewFunctionMap()
			_, err := fm.AddFunction("a", funcType)
			Expect(err).To(BeNil())

			_, err = fm.AddFunction("a", funcType)
			Expect(err).ToNot(BeNil())
		})
		It("Should adjust function definition", func() {
			fm := NewFunctionMap()
			_, err := fm.AddFunction("a", funcType)
			Expect(err).To(BeNil())
			err = fm.AdjustFunction("a", intType)
			Expect(err).To(BeNil())
			Expect(fm.GetFunction("a").GenerateGo()).To(Equal(intType.GenerateGo()))
		})
	})
})

package expressionParsing_test

import (
	. "github.com/apoydence/gof/parser/expressionParsing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("FunctionMap", func() {
	Context("Multiple Functions", func() {
		It("Should return incrementing values", func() {
			fm := NewFunctionMap()
			intType := NewPrimTypeDefinition("int32")
			fi := NewFuncTypeDefinition("", intType, intType)
			a, err := fm.AddFunction("a", fi)
			Expect(err).To(BeNil())
			b, err := fm.AddFunction("b", fi)
			Expect(err).To(BeNil())
			c, err := fm.AddFunction("c", fi)
			Expect(err).To(BeNil())

			Expect(a).To(Equal("a"))
			Expect(b).To(Equal("b"))
			Expect(c).To(Equal("c"))

			Expect(fm.GetFunction("a").(FuncTypeDefinition).FuncName()).To(Equal(a))
			Expect(fm.GetFunction("b").(FuncTypeDefinition).FuncName()).To(Equal(b))
			Expect(fm.GetFunction("c").(FuncTypeDefinition).FuncName()).To(Equal(c))
		})
		It("Should return null for an unknown function", func() {
			fm := NewFunctionMap()
			Expect(fm.GetFunction("x")).To(BeNil())
		})
		It("Should return an error when add the same function name", func() {
			fm := NewFunctionMap()
			intType := NewPrimTypeDefinition("int32")
			fi := NewFuncTypeDefinition("", intType, intType)

			_, err := fm.AddFunction("a", fi)
			Expect(err).To(BeNil())

			_, err = fm.AddFunction("a", fi)
			Expect(err).ToNot(BeNil())
		})
	})
})

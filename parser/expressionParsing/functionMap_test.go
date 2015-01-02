package expressionParsing_test

import (
	. "github.com/apoydence/GoF/parser/expressionParsing"

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

			Expect(a).To(Equal("_0"))
			Expect(b).To(Equal("_1"))
			Expect(c).To(Equal("_2"))

			Expect(fm.GetFunction("a").FuncName()).To(Equal(a))
			Expect(fm.GetFunction("b").FuncName()).To(Equal(b))
			Expect(fm.GetFunction("c").FuncName()).To(Equal(c))
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

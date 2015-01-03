package expressionParsing_test

import (
	. "github.com/apoydence/GoF/parser/expressionParsing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TypeDefinition", func() {
	Context("Function Definition", func() {
		It("Should list arguments and return type", func() {
			i := NewPrimTypeDefinition("int32")
			u := NewPrimTypeDefinition("uint32")
			f := NewFuncTypeDefinition("a", i, u)
			Expect(f.Name()).To(BeEquivalentTo("a int32->uint32"))
		})
	})
	Context("ParseFuncTypeDefinition", func() {
		It("Should return a chained FuncTypeDefinition", func() {
			code := "n int32 -> m int32 -> int32"
			fd, err := ParseFuncTypeDefinition(code)
			Expect(err).To(BeNil())
			Expect(fd).ToNot(BeNil())
			Expect(fd.Argument.Name()).To(BeEquivalentTo("int32"))
			Expect(fd.ArgumentName).To(Equal("n"))
			f2 := fd.ReturnType().(*FuncTypeDefinition)
			Expect(f2.Argument.Name()).To(BeEquivalentTo("int32"))
			Expect(f2.ArgumentName).To(Equal("m"))
			Expect(f2.ReturnType().Name()).To(BeEquivalentTo("int32"))
		})
	})
})

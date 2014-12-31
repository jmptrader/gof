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
			f := NewFuncTypeDefinition(i, u)
			Expect(f.Name()).To(BeEquivalentTo("int32->uint32"))
		})
	})
})

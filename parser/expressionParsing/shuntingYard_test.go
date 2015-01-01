package expressionParsing_test

import (
	. "github.com/apoydence/GoF/parser/expressionParsing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ShuntingYard", func() {
	Context("Without funcions", func() {
		It("Should put the expression into RPN", func() {
			exp := "( 5 + 9 ) * 8 / 3 - 2"
			fm := NewFunctionMap()
			ops, err := ToRpn(exp, fm)
			Expect(err).To(BeNil())
			Expect(ops).To(Equal([]string{"5", "9", "+", "8", "*", "3", "/", "2", "-"}))
		})
	})
	Context("Withfuncions", func() {
		It("Should put the expression into RPN", func() {
			exp := "( 5 + 9 ) * a 100 200 / 3 - 2"
			fm := NewFunctionMap()
			intType := NewPrimTypeDefinition("int32")
			f1 := NewFuncTypeDefinition(intType, intType)
			f2 := NewFuncTypeDefinition(intType, f1)
			_, _ = fm.AddFunction("a", f2)
			ops, err := ToRpn(exp, fm)
			Expect(err).To(BeNil())
			Expect(ops).To(Equal([]string{"5", "9", "+", "100", "200", "a", "*", "3", "/", "2", "-"}))
		})
	})
})

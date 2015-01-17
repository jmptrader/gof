package expressionParsing_test

import (
	. "github.com/apoydence/gof/parser/expressionParsing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ShuntingYard", func() {
	Context("Without funcions", func() {
		It("Should put the expression into RPN", func() {
			exp := "( 5 + 9 ) * 8 / 3 - 2"
			fm := NewFunctionMap()
			ops, err := ToRpn(exp, 9, fm)
			Expect(err).To(BeNil())
			Expect(ops).To(Equal([]string{"5", "9", "+", "8", "*", "3", "/", "2", "-"}))
		})
	})
	Context("Withfuncions", func() {
		It("Should put the expression into RPN", func() {
			exp := "( 5 + 9 ) * a 100 200 / 3 - 2"
			fm := NewFunctionMap()
			f, _, _ := ParseTypeDef("func x int32 -> y int32 -> int32")
			_, _ = fm.AddFunction("a", f)
			ops, err := ToRpn(exp, 9, fm)
			Expect(err).To(BeNil())
			Expect(ops).To(Equal([]string{"5", "9", "+", "100", "200", "a", "*", "3", "/", "2", "-"}))
		})
		It("Should use the definition as a normal number", func() {
			exp := "( 5 + 9 ) * a 100 200 / b - 2"
			fm := NewFunctionMap()
			f, _, _ := ParseTypeDef("func x int32 -> y int32 -> int32")
			intType, _, _ := ParseTypeDef("int32")
			_, _ = fm.AddFunction("a", f)
			_, _ = fm.AddFunction("b", intType)
			ops, err := ToRpn(exp, 9, fm)
			Expect(err).To(BeNil())
			Expect(ops).To(Equal([]string{"5", "9", "+", "100", "200", "a", "*", "b", "/", "2", "-"}))
		})
	})
})

package expressionParsing_test

import (
	"github.com/apoydence/gof/parser"
	. "github.com/apoydence/gof/parser/expressionParsing"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Infix", func() {
	Context("No functions", func() {
		It("It should convert RPN to infix", func() {
			rpn := toRpnValues([]string{"5", "6", "+", "7", "/"})
			fm := NewFunctionMap()
			code, td, err := ToInfix(rpn, fm)
			Expect(err).To(BeNil())
			Expect(td.GenerateGo()).To(Equal("int32"))
			Expect(code).To(Equal("((5+6)/7)"))
		})
	})
	Context("With functions", func() {
		It("Should use a definition as a normal number", func() {
			rpn := toRpnValues([]string{"5", "6", "+", "a:a", "/"})
			fm := NewFunctionMap()
			intType, _, _ := ParseTypeDef("int32")
			fm.AddFunction("a", intType)
			code, td, err := ToInfix(rpn, fm)
			Expect(err).To(BeNil())
			Expect(td.GenerateGo()).To(Equal("int32"))
			Expect(code).To(Equal("((5+6)/a())"))
		})
		It("Should use a function with multiple arguments", func() {
			rpn := toRpnValues([]string{"5", "6", "+", "7", "8", "a", "/"})
			fm := NewFunctionMap()
			f, _, _ := ParseTypeDef("func x int32 -> y int32 -> int32")
			fm.AddFunction("a", f)
			code, td, err := ToInfix(rpn, fm)
			Expect(err).To(BeNil())
			Expect(td.GenerateGo()).To(Equal("int32"))
			Expect(code).To(Equal("((5+6)/a(7)(8))"))
		})
		It("Should use a argument as a normal value", func() {
			rpn := toRpnValues([]string{"5", "6", "+", "a", "/"})
			fm := NewFunctionMap()
			intType, _, _ := ParseTypeDef("int32")
			fm.AddFunction("a", NewArgTypeDefinition(intType))
			code, td, err := ToInfix(rpn, fm)
			Expect(err).To(BeNil())
			Expect(td.GenerateGo()).To(Equal("int32"))
			Expect(code).To(Equal("((5+6)/a)"))
		})
		It("Should use an argument function as a normal value", func() {
			rpn := toRpnValues([]string{"a:arg", "5", "rxF"})
			argF, _, _ := ParseTypeDef("func a int32 -> b int32 -> int32")
			rxF, _, _ := ParseTypeDef("func x func aa int32 -> bb int32 -> int32 -> y int32 -> int32")
			fm := NewFunctionMap()
			fm.AddFunction("arg", argF)
			fm.AddFunction("rxF", rxF)
			code, td, err := ToInfix(rpn, fm)
			Expect(err).To(BeNil())
			Expect(td.GenerateGo()).To(Equal("int32"))
			Expect(code).To(Equal("rxF(arg)(5)"))
		})
		It("Should use an argument function as a normal value", func() {
			rpn := toRpnValues([]string{"5", "6", "+", "a:a", "7", "f", "/"})
			fm := NewFunctionMap()
			funcArg, _, _ := ParseTypeDef("func x int32 -> y int32 -> int32")
			funcType, _, _ := ParseTypeDef("func i func j int32 -> k int32 -> int32 -> m int32 -> int32")
			fm.AddFunction("a", funcArg)
			fm.AddFunction("f", funcType)
			code, td, err := ToInfix(rpn, fm)
			Expect(err).To(BeNil())
			Expect(td.GenerateGo()).To(Equal("int32"))
			Expect(code).To(Equal("((5+6)/f(a)(7))"))
		})
		It("Should use a function's output as an argument as a normal value", func() {
			rpn := toRpnValues([]string{"7", "13", "+", "5", "9", "b", "a", "8", "/", "-"})
			fm := NewFunctionMap()
			funcType, _, _ := ParseTypeDef("func a int32 -> int32")
			twoArgFunc, _, _ := ParseTypeDef("func a int32 -> b int32 -> int32")
			fm.AddFunction("a", twoArgFunc)
			fm.AddFunction("b", funcType)
			code, td, err := ToInfix(rpn, fm)
			Expect(err).To(BeNil())
			Expect(td.GenerateGo()).To(Equal("int32"))
			Expect(code).To(Equal("((7+13)-(a(5)(b(9))/8))"))
		})
	})
})

func toRpnValues(tokens []string) []RpnValue {
	results := make([]RpnValue, 0)
	for _, token := range tokens {
		isArg := strings.HasPrefix(token, "a:")
		if isArg {
			token = token[2:]
		}

		if parser.IsNumber(token) || (parser.ValidFunctionName(token) && isArg) {
			rpn := NewPrimRpnValue(token)
			rpn.Argument = isArg || parser.IsNumber(token)
			results = append(results, rpn)
		} else {
			var prec int
			if parser.ValidFunctionName(token) {
				prec = FuncCall
			} else {
				prec = OpPrec(token)
			}
			results = append(results, NewOpRpnValue(token, prec))
		}
	}
	return results
}

package parser_test

import (
	. "github.com/apoydence/gof/parser"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Parser", func() {
	Context("Tokenize", func() {
		It("Should split up a line based on whitespace", func() {
			code := "\ta b c"
			a, rest := Tokenize(code)
			Expect(a).To(Equal("a"))
			Expect(rest).To(Equal("b c"))
		})
		It("Should return an empty 2nd string when it cant split", func() {
			code := "abc"
			a, rest := Tokenize(code)
			Expect(a).To(Equal("abc"))
			Expect(rest).To(Equal(""))
		})
	})
	Context("Lines", func() {
		It("Should break up all the lines in the block", func() {
			block := "a\nb\nc"
			lines := Lines(block)
			Expect(lines).To(Equal([]string{"a", "b", "c"}))
		})
	})
	Context("FromLines", func() {
		It("Should combine the lines into a block", func() {
			lines := []string{"a", "b", "c"}
			block := FromLines(lines)
			Expect(block).To(Equal("a\nb\nc\n"))
		})
	})
	Context("RemoveTabs", func() {
		It("Should remove the first tab from a series of lines", func() {
			lines := []string{"\ta", "\tb", "\tc"}
			Expect(RemoveTabs(lines)).To(Equal([]string{"a", "b", "c"}))
		})
	})
	Context("IsNumber", func() {
		It("Should be distinguish a float or int", func() {
			i := "324"
			d := "324.2"
			f := "324.2f"
			ui := "324ui"
			l := "324l"
			b := "40b"
			hex := "0x234"
			x := "34x.3"
			y := "34.3.1"
			a := "asd"
			Expect(IsNumber(i)).To(BeTrue())
			Expect(IsNumber(d)).To(BeTrue())
			Expect(IsNumber(f)).To(BeTrue())
			Expect(IsNumber(ui)).To(BeTrue())
			Expect(IsNumber(l)).To(BeTrue())
			Expect(IsNumber(b)).To(BeTrue())
			Expect(IsNumber(hex)).To(BeTrue())
			Expect(IsNumber(x)).To(BeFalse())
			Expect(IsNumber(y)).To(BeFalse())
			Expect(IsNumber(a)).To(BeFalse())
		})
	})
	Context("IsBool", func() {
		It("Should be distinguish a boolean", func() {
			t := "true"
			f := "false"
			x := "3"
			y := "whatever"
			Expect(IsBool(t)).To(BeTrue())
			Expect(IsBool(f)).To(BeTrue())
			Expect(IsBool(x)).To(BeFalse())
			Expect(IsBool(y)).To(BeFalse())
		})
	})

	Context("ValidFunctionName", func() {
		It("Should determine a valid function name", func() {
			a := "a"
			b := "ab"
			c := "ab1"
			d := "ab1_"
			x := "1a"
			y := "_a"
			Expect(ValidFunctionName(a)).To(BeTrue())
			Expect(ValidFunctionName(b)).To(BeTrue())
			Expect(ValidFunctionName(c)).To(BeTrue())
			Expect(ValidFunctionName(d)).To(BeTrue())
			Expect(ValidFunctionName(x)).To(BeFalse())
			Expect(ValidFunctionName(y)).To(BeFalse())
		})
	})

	Context("IsOperator", func() {
		It("Should determine if a value is an operator", func() {
			a := "+"
			b := "-"
			c := "*"
			d := "/"
			x := "a"
			y := "b"
			Expect(IsOperator(a)).To(BeTrue())
			Expect(IsOperator(b)).To(BeTrue())
			Expect(IsOperator(c)).To(BeTrue())
			Expect(IsOperator(d)).To(BeTrue())
			Expect(IsOperator(x)).To(BeFalse())
			Expect(IsOperator(y)).To(BeFalse())
		})
	})

	Context("SplitWhitespace", func() {
		It("Should return the number of words given by n", func() {
			line := "hello world\thow are you"
			words := SplitWhitespace(line, 3)
			Expect(words).To(HaveLen(3))
			Expect(words).To(Equal([]string{"hello", "world", "how are you"}))
		})
	})

	Context("GetFirstToken", func() {
		It("Should return the first word and then the rest of the line", func() {
			line := "hello world\thow are you"
			first, rest := GetFirstToken(line)
			Expect(first).To(Equal("hello"))
			Expect(rest).To(Equal("world\thow are you"))
		})
	})
})

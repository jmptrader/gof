package parser_test

import (
	. "github.com/apoydence/GoF/parser"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Parser", func() {
	Context("Tokenize", func() {
		It("Should split up a line based on whitespace", func() {
			code := "a b c"
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
})

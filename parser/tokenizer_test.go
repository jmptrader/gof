package parser_test

import (
	. "github.com/apoydence/GoF/parser"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Parser", func() {
	Context("Tokenize", func() {
		It("Should break up tokens based on newlines", func() {
			_, tokens, err := Tokenize("		func a -> b int -> int ")
			Expect(err).To(BeNil())
			Expect(tokens).To(Equal([]string{"func", "a", "->", "b", "int", "->", "int"}))
		})

		It("Should give the number of tabs at the beginning of the line", func() {
			tabs, _, err := Tokenize("		func a -> b int -> int")
			Expect(err).To(BeNil())
			Expect(tabs).To(Equal(2))
		})

		It("Should return an error if the line starts with any spaces", func() {
			_, _, err := Tokenize("		 func a -> b int -> int")
			Expect(err).ToNot(BeNil())
		})
	})
})

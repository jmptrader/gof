package parser_test

import (
	. "github.com/apoydence/GoF/parser"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("BlockScanner", func() {
	Context("Scan", func() {
		It("Should break up the blocks", func() {
			code := "func a -> b int -> int\n\tb + 9\n\nfunc b -> c int -> int\n\t c - 4\n\n"
			scanner := NewBlockScanner(strings.NewReader(code))

			Expect(scanner.Scan()).To(BeTrue())
			Expect(scanner.Err()).To(BeNil())
			Expect(scanner.Text()).To(Equal("func a -> b int -> int\n\tb + 9"))

			Expect(scanner.Scan()).To(BeTrue())
			Expect(scanner.Err()).To(BeNil())
			Expect(scanner.Text()).To(Equal("func b -> c int -> int\n\tc - 4"))

			Expect(scanner.Scan()).To(BeFalse())
			Expect(scanner.Err()).To(BeNil())
			Expect(scanner.Text()).To(Equal(""))
		})

		It("Should maintain the inner blocks tabs", func() {
			code := "if a == 4\n\tif b > 3\n\t\t5\n\telse\n\t\t6\nelse\n\t9"
			scanner := NewBlockScanner(strings.NewReader(code))

			Expect(scanner.Scan()).To(BeTrue())
			Expect(scanner.Err()).To(BeNil())
			Expect(scanner.Text()).To(Equal("if a == 4\n\tif b > 3\n\t\t5\n\telse\n\t\t6"))

			Expect(scanner.Scan()).To(BeFalse())
			Expect(scanner.Err()).To(BeNil())
			Expect(scanner.Text()).To(Equal("else\n\t9"))
		})
	})
})

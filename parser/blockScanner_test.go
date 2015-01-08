package parser_test

import (
	. "github.com/apoydence/gof/parser"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("BlockScanner", func() {
	Context("No pre reader", func() {
		It("Should return a single block", func() {
			code := "match\n\tblah\n\tblah"
			scanner := NewBlockScanner(strings.NewReader(code), nopPre)

			Expect(scanner.Scan()).To(BeTrue())
			Expect(scanner.Err()).To(BeNil())
			Expect(scanner.Text()).To(Equal("match\n\tblah\n\tblah"))
		})
		It("Should break up the blocks", func() {
			code := "func a -> b int -> int\n\tb + 9\n\nfunc b -> c int -> int\n\t c - 4\n\n"
			scanner := NewBlockScanner(strings.NewReader(code), nopPre)

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
			scanner := NewBlockScanner(strings.NewReader(code), nopPre)

			Expect(scanner.Scan()).To(BeTrue())
			Expect(scanner.Err()).To(BeNil())
			Expect(scanner.Text()).To(Equal("if a == 4\n\tif b > 3\n\t\t5\n\telse\n\t\t6"))

			Expect(scanner.Scan()).To(BeTrue())
			Expect(scanner.Err()).To(BeNil())
			Expect(scanner.Text()).To(Equal("else\n\t9"))
		})
	})
	Context("Pre-Reader removes tabs", func() {
		It("Should break up the blocks", func() {
			code := "\tfunc a -> b int -> int\n\t\tb + 9\n\nfunc b -> c int -> int\n\t\t c - 4\n\n"
			scanner := NewBlockScanner(strings.NewReader(code), removeTabPre)

			Expect(scanner.Scan()).To(BeTrue())
			Expect(scanner.Err()).To(BeNil())
			Expect(scanner.Text()).To(Equal("func a -> b int -> int\n\tb + 9"))
			Expect(scanner.LineNumber()).To(Equal(0))

			Expect(scanner.Scan()).To(BeTrue())
			Expect(scanner.Err()).To(BeNil())
			Expect(scanner.Text()).To(Equal("func b -> c int -> int\n\tc - 4"))
			Expect(scanner.LineNumber()).To(Equal(3))

			Expect(scanner.Scan()).To(BeFalse())
			Expect(scanner.Err()).To(BeNil())
			Expect(scanner.Text()).To(Equal(""))
		})
	})
})

func nopPre(line string) string {
	return line
}

func removeTabPre(line string) string {
	if len(line) > 0 && line[0] == '\t' {
		return line[1:]
	}
	return line
}

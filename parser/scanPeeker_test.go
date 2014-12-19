package parser_test

import (
	"bufio"
	. "github.com/apoydence/GoF/parser"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ScanPeeker", func() {
	Context("Peek", func() {
		It("Reads from the scanner without removing it", func() {
			s := bufio.NewScanner(strings.NewReader("a\nb\nc"))
			sp := NewScanPeeker(s)

			ok, value := sp.Peek()
			Expect(ok).To(BeTrue())
			Expect(value).To(Equal("a"))

			ok, value = sp.Read()
			Expect(ok).To(BeTrue())
			Expect(value).To(Equal("a"))

			ok, value = sp.Peek()
			Expect(ok).To(BeTrue())
			Expect(value).To(Equal("b"))

			ok, value = sp.Read()
			Expect(ok).To(BeTrue())
			Expect(value).To(Equal("b"))

			ok, value = sp.Peek()
			Expect(ok).To(BeTrue())
			Expect(value).To(Equal("c"))

			ok, value = sp.Read()
			Expect(ok).To(BeTrue())
			Expect(value).To(Equal("c"))
		})
	})
})
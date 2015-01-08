package parser_test

import (
	. "github.com/apoydence/gof/parser"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ScanPeeker", func() {
	Context("Read", func() {
		It("Reads a one liner", func() {
			sp := NewScanPeekerStr("abc")

			ok, value, line := sp.Read()
			Expect(ok).To(BeTrue())
			Expect(value).To(Equal("abc"))
			Expect(line).To(Equal(0))
		})
		It("Reads multiple lines", func() {
			sp := NewScanPeekerStr("a\nb\nc")

			ok, value, line := sp.Read()
			Expect(ok).To(BeTrue())
			Expect(value).To(Equal("a"))
			Expect(line).To(Equal(0))

			ok, value, line = sp.Read()
			Expect(ok).To(BeTrue())
			Expect(value).To(Equal("b"))
			Expect(line).To(Equal(1))

			ok, value, line = sp.Read()
			Expect(ok).To(BeTrue())
			Expect(value).To(Equal("c"))
			Expect(line).To(Equal(2))
		})
	})
	Context("Peek", func() {
		It("Reads from the scanner without removing it", func() {
			sp := NewScanPeekerStr("a\nb\nc")

			ok, value, line := sp.Peek()
			Expect(ok).To(BeTrue())
			Expect(value).To(Equal("a"))
			Expect(line).To(Equal(0))

			ok, value, line = sp.Read()
			Expect(ok).To(BeTrue())
			Expect(value).To(Equal("a"))
			Expect(line).To(Equal(0))

			ok, value, line = sp.Peek()
			Expect(ok).To(BeTrue())
			Expect(value).To(Equal("b"))
			Expect(line).To(Equal(1))

			ok, value, line = sp.Read()
			Expect(ok).To(BeTrue())
			Expect(value).To(Equal("b"))
			Expect(line).To(Equal(1))

			ok, value, line = sp.Read()
			Expect(ok).To(BeTrue())
			Expect(value).To(Equal("c"))
			Expect(line).To(Equal(2))

			ok, _, _ = sp.Read()
			Expect(ok).To(BeFalse())
		})
	})
})

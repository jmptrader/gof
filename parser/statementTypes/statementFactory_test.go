package statementTypes_test

import (
	"github.com/apoydence/GoF/parser"
	. "github.com/apoydence/GoF/parser/statementTypes"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("StatementFactory", func() {
	Context("Next", func() {
		msIf := newMockStatement("if")
		msMatch := newMockStatement("match")

		It("Should return the first statement that matches the block", func() {
			code := strings.NewReader("match\n\tblah\n\tblah\n")
			bs := parser.NewBlockScanner(code, nil)
			sf := NewStatementFactory(bs, msIf, msMatch)

			s := sf.Next()
			Expect(s).ToNot(BeNil())
			Expect(s).To(Equal(msMatch))
		})

		It("Should return nil if it doesn't match a statement", func() {
			code := strings.NewReader("something\n\tblah\n\tblah\n")
			bs := parser.NewBlockScanner(code, nil)
			sf := NewStatementFactory(bs, msIf, msMatch)

			s := sf.Next()
			Expect(s).To(BeNil())
		})

		It("Should be able to read multiple statements", func() {
			code := strings.NewReader("match\n\tblah\n\tblah\nif\n\tfoo")
			bs := parser.NewBlockScanner(code, nil)
			sf := NewStatementFactory(bs, msIf, msMatch)

			s := sf.Next()
			Expect(s).ToNot(BeNil())
			Expect(s).To(Equal(msMatch))

			s = sf.Next()
			Expect(s).ToNot(BeNil())
			Expect(s).To(Equal(msIf))
		})

	})
})

type mockStatement struct {
	startsWith string
}

func newMockStatement(startsWith string) Statement {
	return mockStatement{
		startsWith: startsWith,
	}
}

func (ms mockStatement) Parse(block string, nextBlockScanner *parser.ScanPeeker) Statement {
	if strings.HasPrefix(block, ms.startsWith) {
		return ms
	}

	return nil
}
package statementTypes_test

import (
	"github.com/apoydence/gof/parser"
	"github.com/apoydence/gof/parser/expressionParsing"
	. "github.com/apoydence/gof/parser/statementTypes"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("StatementFactory", func() {
	Context("Read", func() {
		msIf := newMockStatement("if")
		msMatch := newMockStatement("match")

		It("Should return the first statement that matches the block", func() {
			code := strings.NewReader("match\n\tblah\n\tblah\n")
			bs := parser.NewBlockScanner(code, nil)
			bp := parser.NewScanPeeker(bs)
			sf := NewStatementFactory(msIf, msMatch)

			s, err := sf.Read(bp)
			Expect(err).To(BeNil())
			Expect(s).ToNot(BeNil())
			Expect(s).To(Equal(msMatch))
		})

		It("Should return nil if it doesn't match a statement", func() {
			code := strings.NewReader("something\n\tblah\n\tblah\n")
			bs := parser.NewBlockScanner(code, nil)
			bp := parser.NewScanPeeker(bs)
			sf := NewStatementFactory(msIf, msMatch)

			s, err := sf.Read(bp)
			Expect(err).To(BeNil())
			Expect(s).To(BeNil())
		})

		It("Should be able to read multiple statements", func() {
			code := strings.NewReader("match\n\tblah\n\tblah\nif\n\tfoo")
			bs := parser.NewBlockScanner(code, nil)
			bp := parser.NewScanPeeker(bs)
			sf := NewStatementFactory(msIf, msMatch)

			s, err := sf.Read(bp)
			Expect(err).To(BeNil())
			Expect(s).ToNot(BeNil())
			Expect(s).To(Equal(msMatch))

			s, err = sf.Read(bp)
			Expect(err).To(BeNil())
			Expect(s).ToNot(BeNil())
			Expect(s).To(Equal(msIf))
		})

	})
})

type mockStatement struct {
	startsWith string
}

func newMockStatement(startsWith string) StatementParser {
	return mockStatement{
		startsWith: startsWith,
	}
}

func (ms mockStatement) Parse(block string, lineNum int, nextBlockScanner *parser.ScanPeeker, factory *StatementFactory) (Statement, parser.SyntaxError) {
	if strings.HasPrefix(block, ms.startsWith) {
		return ms, nil
	}

	return nil, nil
}

func (ms mockStatement) GenerateGo(fm expressionParsing.FunctionMap) (string, expressionParsing.TypeDefinition, parser.SyntaxError) {
	return "", nil, nil
}

func (ms mockStatement) LineNumber() int {
	return -1
}

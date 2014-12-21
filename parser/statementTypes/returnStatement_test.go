package statementTypes_test

import (
	. "github.com/apoydence/GoF/parser/statementTypes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ReturnStatement", func() {
	var statementParser StatementParser
	var factory *StatementFactory

	BeforeEach(func() {
		statementParser = NewReturnStatementParser()
		factory = NewStatementFactory(statementParser)
	})
	Context("Parse", func() {

		It("Should pick out the statement", func() {
			code := "a + 9"
			r := statementParser.Parse(code, nil, factory)
			Expect(r).ToNot(BeNil())
		})
	})
	Context("GenerateGo", func() {
		It("Should generate the proper Go code with no functions", func() {
			code := "( 1 + 11 ) - 10 / 2"
			r := statementParser.Parse(code, nil, factory)
			genGo, err := r.GenerateGo(nil)
			Expect(err).To(BeNil())
			Expect(genGo).To(Equal("func() int {\n\t((1+11)-(10/2))\n}"))
		})
	})
})

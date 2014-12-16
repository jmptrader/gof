package parser_test

import (
	. "github.com/apoydence/GoF/parser"
	"github.com/apoydence/GoF/parser/statementTypes"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("StatementTypeDecider", func() {
	Context("TypeDecider", func() {
		It("Should pick function if the line starts with the keyword func", func() {
			statement, err := StatementTypeDecider(strings.Fields("func a -> int"))
			Expect(err).To(BeNil())
			Expect(statement.Type()).To(Equal(statementTypes.FunctionType))
		})

		It("Should pick if-statement if the line starts with the keyword if", func() {
			statement, err := StatementTypeDecider(strings.Fields("if something\n\tstatement"))
			Expect(err).To(BeNil())
			Expect(statement.Type()).To(Equal(statementTypes.IfType))
		})

		It("Should pick match-statement if the line starts with the keyword match", func() {
			statement, err := StatementTypeDecider(strings.Fields("match something\n\t0 -> statement\n\tn -> default"))
			Expect(err).To(BeNil())
			Expect(statement.Type()).To(Equal(statementTypes.MatchType))
		})

		It("Should pick declaration-statement if the line starts with a declartion", func() {
			statement, err := StatementTypeDecider(strings.Fields("x = func a int -> int"))
			Expect(err).To(BeNil())
			Expect(statement.Type()).To(Equal(statementTypes.DeclarationType))
		})
	})

})

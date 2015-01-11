package statementTypes_test

import (
	"github.com/apoydence/gof/parser/expressionParsing"
	. "github.com/apoydence/gof/parser/statementTypes"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("LambdaStatement", func() {
	var statementParser StatementParser
	var factory *StatementFactory
	BeforeEach(func() {
		retParser := NewReturnStatementParser()
		declParser := NewLetParser()
		statementParser = NewLambdaStatementParser()
		factory = NewStatementFactory(declParser, statementParser, retParser)
	})
	Context("Parse", func() {
		It("Should distinguish a lambda statement", func() {
			code := "func a A -> b B -> int32 ->\n\t5+9"
			f, err := statementParser.Parse(code, 0, nil, factory)
			Expect(f).ToNot(BeNil())
			fs := f.(*LambdaStatement)
			Expect(err).To(BeNil())
			Expect(fs).ToNot(BeNil())
			Expect(fs.TypeDef.Name()).To(BeEquivalentTo("a A->b B->int32"))
			Expect(fs.InnerStatements).ToNot(BeNil())
			Expect(fs.InnerStatements).To(HaveLen(1))
		})
		It("Should distinguish a single line lambda statement", func() {
			code := "func a A -> b B -> int32 -> 5+9"
			f, err := statementParser.Parse(code, 0, nil, factory)
			Expect(f).ToNot(BeNil())
			fs := f.(*LambdaStatement)
			Expect(err).To(BeNil())
			Expect(fs).ToNot(BeNil())
			Expect(fs.TypeDef.Name()).To(BeEquivalentTo("a A->b B->int32"))
			Expect(fs.InnerStatements).ToNot(BeNil())
			Expect(fs.InnerStatements).To(HaveLen(1))
		})
	})
	Context("GenerateGo", func() {
		It("Should generate a proper Go function with one inner statement", func() {
			fm := expressionParsing.NewFunctionMap()
			code := "func a int32 -> b int32 -> c int32 -> int32 -> \n\ta + b + c"
			f, err := statementParser.Parse(code, 0, nil, factory)
			Expect(f).ToNot(BeNil())
			Expect(err).To(BeNil())
			code, _, err = f.GenerateGo(fm)
			Expect(err).To(BeNil())
			Expect(matchCode(code, "func (a int32) func (b int32) func (c int32) int32{\n\treturn func (b int32) func (c int32) int32 {\n\t\treturn func (c int32) int32 {\n\t\t\t\n\t\t\treturn ((a+b)+c)\n\t\t}\n\t}\n}")).To(BeTrue())
		})
		It("Should generate a proper Go package function with one inner statement", func() {
			fm := expressionParsing.NewFunctionMap()
			code := "func a int32 -> b int32 -> c int32 -> int32 -> \n\tlet x = func z int32 -> int32 -> \n\t\tz + 1\n\ta + b + x c"
			retParser := NewReturnStatementParser()
			letParser := NewLetParser()
			statementParser := NewPackageLambdaStatementParser("add")
			packFactory := NewStatementFactory(letParser, statementParser, retParser)
			f, err := statementParser.Parse(code, 0, nil, packFactory)
			Expect(f).ToNot(BeNil())
			Expect(err).To(BeNil())
			code, _, err = f.GenerateGo(fm)
			Expect(err).To(BeNil())
			Expect(matchCode(code, "func add (a int32) func (b int32) func (c int32) int32{\n\treturn func (b int32) func (c int32) int32 {\n\t\treturn func (c int32) int32 {\n\t\t\t\n\t\t\tvar x func(z int32) int32\nx = func (z int32) int32{\nreturn (z+1)\n}\nreturn ((a+b)+x(c))\n\t\t}\n\t}\n}")).To(BeTrue())
		})
		It("Should generate a proper Go function with multiple inner statements", func() {
			fm := expressionParsing.NewFunctionMap()
			code := "func a int32 -> b int32 -> c int32 -> int32 ->\n\tlet d = 6\n\ta + b + c + d"
			f, err := statementParser.Parse(code, 0, nil, factory)
			Expect(f).ToNot(BeNil())
			Expect(err).To(BeNil())
			code, _, err = f.GenerateGo(fm)
			Expect(err).To(BeNil())
			Expect(matchCode(code, "func (a int32) func (b int32) func (c int32) int32{\n\treturn func (b int32) func (c int32) int32 {\n\t\treturn func (c int32) int32 {\n\t\t\t\n\t\t\tvar d func() int32\nd = func(){return 6}return (((a+b)+c)+d())\n\t\t}\n\t}\n}")).To(BeTrue())
		})
	})
	Context("Proper statement structure", func() {
		It("Should end in a return statement", func() {
			code := "func a int32 -> b int32 -> c int32 -> int32 ->\n\tlet d = 6"
			_, err := statementParser.Parse(code, 0, nil, factory)
			Expect(err).ToNot(BeNil())
		})
		It("Should only have declaration statement at the beginning", func() {
			code := "func a int32 -> b int32 -> c int32 -> int32 ->\n\ta + 6\n\tb + c"
			_, err := statementParser.Parse(code, 0, nil, factory)
			Expect(err).ToNot(BeNil())
		})
	})
})

func matchCode(codeA, codeB string) bool {
	return stripWhitespace(codeA) == stripWhitespace(codeB)
}

func stripWhitespace(s string) string {
	noTabs := strings.Replace(s, "\t", "", -1)
	noSpaces := strings.Replace(noTabs, " ", "", -1)
	return strings.Replace(noSpaces, "\n", "", -1)
}

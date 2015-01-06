package generate

import (
	"github.com/apoydence/gof/parser"
	"github.com/apoydence/gof/parser/expressionParsing"
	"github.com/apoydence/gof/parser/statementTypes"
	"io"
)

func GofToGo(reader io.Reader, writer io.Writer) error {
	bs := parser.NewBlockScanner(reader, nil)
	bp := parser.NewScanPeeker(bs)
	factory := statementTypes.NewStatementFactory(fetchStatementParsers()...)
	funcMap := expressionParsing.NewFunctionMap()

	io.WriteString(writer, "package tests\n\n")

	return writeGeneratedBlocks(factory, bp, funcMap, writer)
}

func writeGeneratedBlocks(factory *statementTypes.StatementFactory, peeker *parser.ScanPeeker, fm expressionParsing.FunctionMap, writer io.Writer) parser.SyntaxError {
	s, synErr := factory.Read(peeker)
	if synErr != nil {
		return synErr
	}

	if s != nil {
		code, _, synErr := s.GenerateGo(fm)
		if synErr != nil {
			return synErr
		}

		_, err := io.WriteString(writer, code+"\n\n")

		if err != nil {
			return parser.NewSyntaxError(err.Error(), 0, 0)
		}
		return writeGeneratedBlocks(factory, peeker, fm, writer)
	}

	return nil
}

func fetchStatementParsers() []statementTypes.StatementParser {
	s := make([]statementTypes.StatementParser, 0)
	s = append(s, statementTypes.NewDeclarationParser())
	s = append(s, statementTypes.NewFunctionStatementParser())
	s = append(s, statementTypes.NewReturnStatementParser())
	return s
}

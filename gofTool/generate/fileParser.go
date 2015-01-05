package generate

import (
	"github.com/apoydence/GoF/parser"
	"github.com/apoydence/GoF/parser/expressionParsing"
	"github.com/apoydence/GoF/parser/statementTypes"
	"io"
)

func GofToGo(reader io.Reader, writer io.Writer) error {
	bs := parser.NewBlockScanner(reader, nil)
	bp := parser.NewScanPeeker(bs)
	factory := statementTypes.NewStatementFactory(fetchStatementParsers()...)
	funcMap := expressionParsing.NewFunctionMap()

	io.WriteString(writer, "package gof\n\n")

	return writeGeneratedBlocks(factory, bp, funcMap, writer)
}

func writeGeneratedBlocks(factory *statementTypes.StatementFactory, peeker *parser.ScanPeeker, fm expressionParsing.FunctionMap, writer io.Writer) error {
	if s := factory.Read(peeker); s != nil {
		code, _, err := s.GenerateGo(fm)
		if err != nil {
			return err
		}

		_, err = io.WriteString(writer, code)

		if err != nil {
			return err
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

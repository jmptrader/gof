package statementTypes

import "github.com/apoydence/GoF/parser"

type Statement interface {
	Parse(block string, nextBlockScanner *parser.ScanPeeker) Statement
}

type StatementFactory struct {
	blockPeeker *parser.ScanPeeker
	statements  []Statement
}

func NewStatementFactory(blockScanner *parser.BlockScanner, statements ...Statement) *StatementFactory {
	return &StatementFactory{
		blockPeeker: parser.NewScanPeeker(blockScanner),
		statements:  statements,
	}
}

func (sf *StatementFactory) Next() Statement {
	ok, value := sf.blockPeeker.Read()

	if !ok {
		return nil
	}

	for _, s := range sf.statements {
		statement := s.Parse(value, sf.blockPeeker)
		if statement != nil {
			return statement
		}
	}

	return nil
}

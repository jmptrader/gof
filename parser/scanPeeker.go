package parser

import (
	"strings"
)

type Scanner interface {
	Scan() bool
	Text() string
	Err() error
	LineNumber() int
}

type ScanPeeker struct {
	scanner    Scanner
	peeked     string
	readPeeked bool
	scanCheck  bool
	line       int
}

func NewScanPeeker(scanner Scanner) *ScanPeeker {
	return &ScanPeeker{
		scanner:   scanner,
		scanCheck: true,
	}
}

func NewScanPeekerStr(block string) *ScanPeeker {
	return NewScanPeeker(NewBlockScanner(strings.NewReader(block), nil))
}

func (sp *ScanPeeker) Peek() (bool, string, int) {
	ok, value, line := sp.Read()
	sp.peeked = value
	sp.readPeeked = true
	sp.line = line
	return ok, sp.peeked, sp.line
}

func (sp *ScanPeeker) Read() (bool, string, int) {
	defer func() {
		sp.readPeeked = false
	}()

	if sp.readPeeked {
		sp.readPeeked = false
		return true, sp.peeked, sp.line
	}

	if !sp.scanCheck {
		return false, "", -1
	}

	sp.scanCheck = sp.scanner.Scan()
	return sp.scanCheck && sp.scanner.Err() == nil, sp.scanner.Text(), sp.scanner.LineNumber()
}

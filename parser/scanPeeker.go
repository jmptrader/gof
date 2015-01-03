package parser

import (
	"bufio"
	"strings"
)

type Scanner interface {
	Scan() bool
	Text() string
	Err() error
}

type ScanPeeker struct {
	scanner    Scanner
	peeked     string
	readPeeked bool
	scanCheck  bool
}

func NewScanPeeker(scanner Scanner) *ScanPeeker {
	return &ScanPeeker{
		scanner:   scanner,
		scanCheck: true,
	}
}

func NewScanPeekerStr(block string) *ScanPeeker {
	return NewScanPeeker(bufio.NewScanner(strings.NewReader(block)))
}

func (sp *ScanPeeker) Peek() (bool, string) {
	ok, value := sp.Read()
	sp.peeked = value
	sp.readPeeked = true
	return ok, sp.peeked
}

func (sp *ScanPeeker) Read() (bool, string) {
	defer func() {
		sp.readPeeked = false
	}()

	if sp.readPeeked {
		sp.readPeeked = false
		return true, sp.peeked
	}

	if !sp.scanCheck {
		return false, ""
	}

	sp.scanCheck = sp.scanner.Scan()
	return sp.scanCheck && sp.scanner.Err() == nil, sp.scanner.Text()
}

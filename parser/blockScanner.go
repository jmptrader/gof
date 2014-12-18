package parser

import (
	"bufio"
	"io"
	"strings"
)

type BlockScanner struct {
	err      error
	block    string
	prevLine string
	scanner  *bufio.Scanner
}

func NewBlockScanner(code io.Reader) *BlockScanner {
	return &BlockScanner{
		scanner: bufio.NewScanner(code),
	}
}

func (bs *BlockScanner) Scan() bool {
	bs.block = ""
	enteredBlock := false
	for len(bs.prevLine) > 0 || bs.scanner.Scan() {
		var text string
		if len(bs.prevLine) > 0 {
			text = bs.prevLine
		} else {
			text = bs.scanner.Text()
		}
		tabs := startsWithTab(text)
		tabbed := tabs > 0
		if (!enteredBlock && !tabbed) || (enteredBlock && tabbed) {
			enteredBlock = true
			bs.block = bs.block + "\n" + trim(text, tabs)
			bs.prevLine = ""
		} else {
			bs.prevLine = bs.scanner.Text()
			return true
		}
	}

	bs.err = bs.scanner.Err()
	return false
}

func trim(line string, tabs int) string {
	return strings.Repeat("\t", tabs) + strings.TrimSpace(line)
}

func startsWithTab(line string) int {
	tabs := 0
	for _, c := range line {
		if c == '\t' {
			tabs++
		} else {
			break
		}
	}
	return tabs
}

func (bs *BlockScanner) Err() error {
	return bs.err
}

func (bs *BlockScanner) Text() string {
	return strings.TrimSpace(bs.block)
}

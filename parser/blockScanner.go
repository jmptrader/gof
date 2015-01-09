package parser

import (
	"bufio"
	"errors"
	"io"
	"strings"
)

type BlockScanner struct {
	err       error
	block     string
	prevLine  string
	scanner   *bufio.Scanner
	pre       func(string) string
	line      int
	blockLine int
}

func NewBlockScanner(code io.Reader, preRead func(string) string) *BlockScanner {
	if preRead == nil {
		preRead = func(s string) string {
			return s
		}
	}

	return &BlockScanner{
		scanner: bufio.NewScanner(code),
		pre:     preRead,
	}
}

func (bs *BlockScanner) Scan() bool {
	bs.block = ""
	enteredBlock := false
	removeTabs := 0
	firstLine := true
	prevTab := 0
	for len(bs.prevLine) > 0 || bs.scanner.Scan() {
		var text string
		if len(bs.prevLine) > 0 {
			text = bs.prevLine
			bs.blockLine--
		} else {
			text = bs.getLineFromScanner()
			bs.line++
		}

		tabs := startsWithTab(text)

		if tabs > prevTab+1 {
			bs.err = errors.New("Incorrect number of tabs")
			break
		}

		prevTab = tabs

		if firstLine && len(strings.TrimSpace(text)) > 0 {
			bs.blockLine = bs.line - 1
			firstLine = false
			if tabs != 0 {
				bs.err = errors.New("Incorrect number of tabs")
				break
			}
		}

		if !enteredBlock && tabs > 0 || removeTabs > 0 && len(text) > 0 && text[0] == '\t' {
			removeTabs = tabs
			text = text[1:]
			tabs--
		}

		tabbed := tabs > 0

		if (!enteredBlock && !tabbed) || (enteredBlock && tabbed) {
			if len(strings.TrimSpace(text)) > 0 {
				enteredBlock = true
				bs.block = bs.block + "\n" + trim(text, tabs)
				bs.prevLine = ""
			}
		} else {
			bs.prevLine = text
			enteredBlock = false
			return true
		}
	}

	if bs.err == nil {
		bs.err = bs.scanner.Err()
	}
	return bs.err == nil && enteredBlock
}

func (bs *BlockScanner) getLineFromScanner() string {
	return bs.pre(bs.scanner.Text())
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

func (bs *BlockScanner) LineNumber() int {
	return bs.blockLine
}

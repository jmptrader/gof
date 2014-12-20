package parser

import (
	"bufio"
	"regexp"
	"strings"
)

var reg *regexp.Regexp

func init() {
	reg = regexp.MustCompile("\\s")
}

func Tokenize(line string) (string, string) {
	values := reg.Split(line, 2)
	if len(values) == 2 {
		return values[0], values[1]
	}
	return values[0], ""
}

func Lines(block string) []string {
	scanner := bufio.NewScanner(strings.NewReader(block))
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

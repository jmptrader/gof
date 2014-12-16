package parser

import (
	"errors"
	"strings"
)

func Tokenize(line string) (int, []string, error) {
	tabs, trimmedLine := splitTabs(line)
	if trimmedLine[0] == ' ' {
		return -1, nil, errors.New("Lines can not begin with a space. Use tabs.")
	}

	return tabs, strings.Fields(trimmedLine), nil
}

func splitTabs(line string) (int, string) {
	tabs := 0
	for _, c := range line {
		if c == '	' {
			tabs++
		} else {
			break
		}
	}
	return tabs, line[tabs:]
}

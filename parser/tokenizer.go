package parser

import (
	"bufio"
	"regexp"
	"strings"
)

var regToken *regexp.Regexp
var regNum *regexp.Regexp
var regFunc *regexp.Regexp
var regOps *regexp.Regexp
var ops []string = []string{"+", "-", "*", "/"}

func init() {
	regToken = regexp.MustCompile("\\s")
	regNum = regexp.MustCompile("^(0x)?[0-9]+((u?[bhl])|(ui)|f|(\\.[0-9]+f?))?$")
	regFunc = regexp.MustCompile("^[a-zA-Z][a-zA-z0-9]*$")
	regOps = regexp.MustCompile("^" + buildOpsPattern(ops) + "$")
}

func Tokenize(line string) (string, string) {
	values := regToken.Split(line, 2)
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

func IsNumber(value string) bool {
	return regNum.MatchString(value)
}

func IsBool(value string) bool {
	return value == "true" || value == "false"
}

func IsPrimitive(value string) bool {
	return IsNumber(value) || IsBool(value)
}

func ValidFunctionName(value string) bool {
	return regFunc.MatchString(value)
}

func IsOperator(value string) bool {
	return regOps.MatchString(value)
}

func buildOpsPattern(ops []string) string {
	result := "["
	for _, o := range ops {
		result = result + "\\" + o
	}
	return result + "]"
}

package expressionParsing_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestExpressionParsing(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ExpressionParsing Suite")
}

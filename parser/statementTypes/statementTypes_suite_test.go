package statementTypes_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestStatementTypes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "StatementTypes Suite")
}

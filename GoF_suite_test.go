package GoF_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGoF(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoF Suite")
}

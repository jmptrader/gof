package tests_test

import (
	. "github.com/apoydence/gof/tool/tests"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("FuncAsArg", func() {
	Context("Has an inner function that has an argument that is a function", func() {
		It("Should calculate the correct value", func() {
			result := FuncAsArg(2)(3)
			Expect(result).To(BeEquivalentTo(16))
		})
	})
})

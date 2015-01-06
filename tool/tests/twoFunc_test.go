package tests_test

import (
	. "github.com/apoydence/GoF/tool/tests"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TwoFunc", func() {
	Context("Action", func() {
		It("Should multiply the first two values and then add the third", func() {
			Expect(Action(2)(3)(4)).To(BeEquivalentTo(10))
		})
	})
})

package tests_test

import (
	. "github.com/apoydence/gof/tool/tests"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TwoFunc", func() {
	Context("Action", func() {
		It("Should multiply the first two values and then add the third", func() {
			Expect(Action(2)(3)(4)).To(BeEquivalentTo(15))
			Expect(Action(3)(4)(5)).To(BeEquivalentTo(24))
		})
	})
})

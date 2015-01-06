package tests_test

import (
	. "github.com/apoydence/GoF/tool/tests"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SingleFunc", func() {
	Context("AddTogether", func() {
		It("Should add the three numbers together", func() {
			Expect(AddTogether(1)(2)(3)).To(BeEquivalentTo(6))
			addTo5 := AddTogether(5)
			Expect(addTo5(1)(2)).To(BeEquivalentTo(8))
			addTo5And6 := addTo5(6)
			Expect(addTo5And6(1)).To(BeEquivalentTo(12))
		})
	})
})

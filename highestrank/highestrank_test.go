package highestrank_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/mentalclear-learningo/katas/highestrank"
)

var _ = Describe("Tests", func() {
	Describe("Sample tests", func() {
		It("Sample test 1: 12, 10, 8, 12, 7, 6, 4, 10, 12", func() {
			Expect(HighestRankRef([]int{12, 10, 8, 12, 7, 6, 4, 10, 12})).To(Equal(12))
		})
	})
})

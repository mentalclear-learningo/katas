package zerostoend_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "zerostoend"
)

var _ = Describe("Zerostoend", func() {
	It("should test that the solution returns the correct value", func() {
		Expect(MoveZeros([]int{1, 2, 0, 1, 0, 1, 0, 3, 0, 1})).To(Equal([]int{1, 2, 1, 1, 3, 1, 0, 0, 0, 0}))
	})
})

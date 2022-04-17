package finduniq_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/mentalclear-learningo/katas/finduniq"
)

var _ = Describe("Finduniq", func() {
	It("Basic Case 1", func() {
		Expect(FindUniq([]float32{1.0, 1.0, 1.0, 2.0, 1.0, 1.0})).To(Equal(float32(2)))
	})
	It("Basic Case 2", func() {
		Expect(FindUniq([]float32{0, 0, 0.55, 0, 0})).To(Equal(float32(0.55)))
	})
})

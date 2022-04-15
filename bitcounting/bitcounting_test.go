package bitcounting_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/mentalclear-learningo/katas/bitcounting"
)

var _ = Describe("Bitcounting", func() {
	It("Test for 0", func() {
		Expect(CountBits(0)).To(Equal(0))
	})
	It("Test for 4", func() {
		Expect(CountBits(4)).To(Equal(1))
	})
	It("Test for 7", func() {
		Expect(CountBits(7)).To(Equal(3))
	})
	It("Test for 9", func() {
		Expect(CountBits(9)).To(Equal(2))
	})
	It("Test for 10", func() {
		Expect(CountBits(10)).To(Equal(2))
	})
})

var _ = Describe("Bitcounting Simplified", func() {
	It("Test for 0", func() {
		Expect(CountBitsSimplified(0)).To(Equal(0))
	})
	It("Test for 4", func() {
		Expect(CountBitsSimplified(4)).To(Equal(1))
	})
	It("Test for 7", func() {
		Expect(CountBitsSimplified(7)).To(Equal(3))
	})
	It("Test for 9", func() {
		Expect(CountBitsSimplified(9)).To(Equal(2))
	})
	It("Test for 10", func() {
		Expect(CountBitsSimplified(10)).To(Equal(2))
	})
})

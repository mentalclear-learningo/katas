package createphone_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "createphone"
)

var _ = Describe("Createphone", func() {
	It("Test CreatePhoneNumber func", func() {
		Expect(CreatePhoneNumber([10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0})).To(Equal("(123) 456-7890"))
	})
	It("Test CreatePhoneNumberRefactored func", func() {
		Expect(CreatePhoneNumberRefactored([10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0})).To(Equal("(123) 456-7890"))
	})
})

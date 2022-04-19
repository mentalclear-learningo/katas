package weirdstring_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/mentalclear-learningo/katas/weirdstring"
)

var _ = Describe("Sample Test Cases:", func() {
	It("Should return the correct values", func() {
		Expect(ToWeirdCase("abc def")).To(Equal("AbC DeF"))
	})

	It("Should return the correct values", func() {
		Expect(ToWeirdCase("ABC")).To(Equal("AbC"))
	})

	It("Should return the correct values", func() {
		Expect(ToWeirdCase("This is a test Looks like you passed")).To(Equal("ThIs Is A TeSt LoOkS LiKe YoU PaSsEd"))
	})
})

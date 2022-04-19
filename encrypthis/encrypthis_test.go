package encrypthis_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/mentalclear-learningo/katas/encrypthis"
)

var _ = Describe("Test Suite", func() {
	It("Sample Tests", func() {
		Expect(EncryptThis("")).Should(Equal(""))
	})
	It("Sample Tests", func() {
		Expect(EncryptThis("A wise old owl lived in an oak")).Should(Equal("65 119esi 111dl 111lw 108dvei 105n 97n 111ka"))
	})
	It("Sample Tests", func() {
		Expect(EncryptThis("The more he saw the less he spoke")).Should(Equal("84eh 109ero 104e 115wa 116eh 108sse 104e 115eokp"))
	})
	It("Sample Tests", func() {
		Expect(EncryptThis("The less he spoke the more he heard")).Should(Equal("84eh 108sse 104e 115eokp 116eh 109ero 104e 104dare"))
	})
	It("Sample Tests", func() {
		Expect(EncryptThis("Why can we not all be like that wise old bird")).Should(Equal("87yh 99na 119e 110to 97ll 98e 108eki 116tah 119esi 111dl 98dri"))
	})
	It("Sample Tests", func() {
		Expect(EncryptThis("Thank you Piotr for all your help")).Should(Equal("84kanh 121uo 80roti 102ro 97ll 121ruo 104ple"))
	})
})

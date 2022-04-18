package whichin_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/mentalclear-learningo/katas/whichin"
)

func dotest(array1 []string, array2 []string, exp []string) {
	var ans = InArray(array1, array2)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Tests Tour", func() {

	It("Basic Test Case 1", func() {
		var a1 = []string{"live", "arp", "strong"}
		var a2 = []string{"lively", "alive", "harp", "sharp", "armstrong"}
		var r = []string{"arp", "live", "strong"}
		dotest(a1, a2, r)
	})

	It("Basic Test Case 2", func() {
		var a1 = []string{"cod", "code", "wars", "ewar", "ar"}
		var a2 = []string{}
		var r = []string{}
		dotest(a1, a2, r)
	})

	It("Basic Test Case 3", func() {
		var a1 = []string{"duplicates", "duplicates"}
		var a2 = []string{"duplicates", "duplicates"}
		var r = []string{"duplicates"}
		dotest(a1, a2, r)
	})

	It("Basic Test Case 4", func() {
		var a1 = []string{"ou", "ple", "wh", "or", "oesi", "pini", "he", "ing", "ve", "ing", "omm", "byc", "vec", "tor", "ouv", "by", "oes"}
		var a2 = []string{"answer", "1.9", "somewhere", "does", "ruby-doc", "most", "here", "I", "opinion", "comment", "out", "questions", "have", "to", "have", "perfect", "I", "versioning", "a", "1.9.2.", "is", "using", "your", "have", "would", "reference", "your", "does", "sample", "for", "I", "Ruby"}
		var r = []string{"by", "he", "ing", "oes", "omm", "or", "ou", "pini", "ple", "ve", "wh"}
		dotest(a1, a2, r)
	})
})

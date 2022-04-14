package multsthreefive_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "multsthreefive"
)

var _ = Describe("Mults", func() {
	It("should handle basic cases", func() {
		Expect(Multiple3And5(10)).To(Equal(23))
	})
})

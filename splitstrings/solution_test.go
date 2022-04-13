package splitstings_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "splitstrings"
)

var _ = Describe("Solution", func() {
	It("Basic tests", func() {
		Expect(Solution("abc")).To(Equal([]string{"ab", "c_"}))
		Expect(Solution("dawsd")).To(Equal([]string{"da", "ws", "d_"}))
		Expect(Solution("awsaws")).To(Equal([]string{"aw", "sa", "ws"}))
	})
})

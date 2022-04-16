package josephusrv_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/mentalclear-learningo/katas/josephusrv"
)

var _ = Describe("Sample Tests", func() {
	It("should handle basic tests", func() {
		dotest(7, 3, 4)
	})
	// It("should handle basic tests", func() {
	// 	dotest(11, 19, 10)
	// })
	// It("should handle basic tests", func() {
	// 	dotest(40, 3, 28)
	// })
	// It("should handle basic tests", func() {
	// 	dotest(14, 2, 13)
	// })
	// It("should handle basic tests", func() {
	// 	dotest(100, 1, 100)
	// })
})

func dotest(n, k, e int) {
	fmt.Println(n, k)
	Expect(JosephusSurvivor(n, k)).To(Equal(e))
}

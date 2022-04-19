package ipaddr_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/mentalclear-learningo/katas/ipaddr"
)

var _ = Describe("Test Example", func() {
	It("should test that 12.255.56.1 is correct", func() {
		Expect(Is_valid_ip("12.255.56.1")).To(Equal(true))
	})
})

var _ = Describe("Test Example", func() {
	It("should test that '' is incorrect", func() {
		Expect(Is_valid_ip("")).To(Equal(false))
	})
})

var _ = Describe("Test Example", func() {
	It("should test that abc.def.ghi.jkl is incorrect", func() {
		Expect(Is_valid_ip("abc.def.ghi.jkl")).To(Equal(false))
	})
})

var _ = Describe("Test Example", func() {
	It("should test that 123.456.789.0 is incorrect", func() {
		Expect(Is_valid_ip("123.456.789.0")).To(Equal(false))
	})
})

var _ = Describe("Test Example", func() {
	It("should test that 12.34.56 is incorrect", func() {
		Expect(Is_valid_ip("12.34.56")).To(Equal(false))
	})
})

var _ = Describe("Test Example", func() {
	It("should test that 12.34.56 .1 is incorrect", func() {
		Expect(Is_valid_ip("12.34.56 .1")).To(Equal(false))
	})
})

var _ = Describe("Test Example", func() {
	It("should test that 12.34.56.-1 is incorrect", func() {
		Expect(Is_valid_ip("12.34.56.-1")).To(Equal(false))
	})
})

var _ = Describe("Test Example", func() {
	It("should test that 123.045.067.089 is incorrect", func() {
		Expect(Is_valid_ip("123.045.067.089")).To(Equal(false))
	})
})

var _ = Describe("Test Example", func() {
	It("should test that 127.1.1.0 is correct", func() {
		Expect(Is_valid_ip("127.1.1.0")).To(Equal(true))
	})
})

var _ = Describe("Test Example", func() {
	It("should test that 0.0.0.0 is correct", func() {
		Expect(Is_valid_ip("0.0.0.0")).To(Equal(true))
	})
})

var _ = Describe("Test Example", func() {
	It("should test that 0.34.82.53 is correct", func() {
		Expect(Is_valid_ip("0.34.82.53")).To(Equal(true))
	})
})

var _ = Describe("Test Example", func() {
	It("should test that 192.168.1.300 is incorrect", func() {
		Expect(Is_valid_ip("192.168.1.300")).To(Equal(false))
	})
})

package bitcounting_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestBitcounting(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Bitcounting Suite")
}

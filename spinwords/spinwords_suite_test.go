package spinwords_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestSpinwords(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Spinwords Suite")
}

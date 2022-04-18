package mexwave_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestMexwave(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Mexwave Suite")
}

package oddint_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestOddint(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Oddint Suite")
}

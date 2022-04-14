package multsthreefive_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestMultsthreefive(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Multsthreefive Suite")
}

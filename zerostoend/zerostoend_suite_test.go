package zerostoend_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestZerostoend(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Zerostoend Suite")
}

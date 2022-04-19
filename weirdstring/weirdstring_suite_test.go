package weirdstring_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestWeirdstring(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Weirdstring Suite")
}

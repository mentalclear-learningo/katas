package deadfish_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestDeadfish(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Deadfish Suite")
}

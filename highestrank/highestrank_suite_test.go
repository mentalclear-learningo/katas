package highestrank_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestHighestrank(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Highestrank Suite")
}

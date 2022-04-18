package whichin_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestWhichin(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Whichin Suite")
}

package encrypthis_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestEncrypthis(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Encrypthis Suite")
}

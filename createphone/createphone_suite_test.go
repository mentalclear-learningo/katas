package createphone_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCreatephone(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Createphone Suite")
}

package josephusrv_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestJosephusrv(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Josephusrv Suite")
}

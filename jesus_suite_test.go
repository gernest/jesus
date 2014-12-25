package jesus_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestJesus(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Jesus Suite")
}

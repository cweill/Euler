package problem42_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestProblem42(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Problem42 Suite")
}

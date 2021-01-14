package one_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestOrder(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "One suite")
}

var _ = Describe("as a user with random input", func() {
	Describe("and last value odd", func() {
		It("should return a bigger value than a even run", func() {
			//Expect(true).To(Equal(false))
		})
	})
})

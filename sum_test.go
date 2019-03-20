package main

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestApp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "main")
}

var _ = Describe("Test main", func() {
	Context("Adding two values", func() {
		It("should return 10", func() {
			Expect(sum(5, 5)).To(Equal(10))
			Expect(sum(0, 5)).To(Equal(5))
		})
	})
	Context("Input", func() {
		It("when valid input should return two int's", func() {
			input := []string{"A", "1", "2"}
			a, b, err := getInput(input)
			Expect(err).To(BeNil())
			Expect(a).To(Equal(1))
			Expect(b).To(Equal(2))
		})
		It("when input array doest not contain all elements", func() {
			input := []string{}
			_, _, err := getInput(input)
			Expect(err).To(HaveOccurred())
		})
		It("when input array has string", func() {
			input := []string{"S", "MC", "YX"}
			_, _, err := getInput(input)
			Expect(err).To(HaveOccurred())
		})
	})
})

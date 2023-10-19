package main

import (
	"math"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCalculator(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Calculator Suite")
}

var _ = Describe("Adding functionality", func() {
	Context("Addition of two numbers", func() {
		It("Should return 4 when 2 and 2 are added", func() {
			Expect(Add(2, 2)).To(Equal(4))
		})
	})

	Context("Addition of two numbers", func() {
		It("Should return 5 when 3 and 2 are added", func() {
			Expect(Add(3, 2)).To(Equal(5))
		})
	})
})

var _ = Describe("Divide functionality", func() {
	It("Should return 2 when 4 is divided by 2", func() {
		Expect(Divide(4, 2)).To(Equal(2))
	})

	It("Should return 0 when 0 is divided by 2", func() {
		Expect(Divide(0, 2)).To(Equal(0))
	})

	It("Should return -1 when 2 is divided by 0", func() {
		Expect(Divide(2, 0)).To(Equal(math.MinInt))
	})
})

var _ = Describe("Multiply functionality", func() {
	It("Should return 4 when 2 is multiplied by 2", func() {
		Expect(Multiply(2, 2)).To(Equal(4))
	})

	It("Should return 0 when 0 is multiplied by 2", func() {
		Expect(Multiply(0, 2)).To(Equal(0))
	})

	It("Should return 0 when 2 is multiplied by 0", func() {
		Expect(Multiply(2, 0)).To(Equal(0))
	})
})

var _ = Describe("Subtract functionality", func() {
	It("Should return 0 when 2 is subtracted from 2", func() {
		Expect(Subtract(2, 2)).To(Equal(0))
	})

	It("Should return 2 when 0 is subtracted from 2", func() {
		Expect(Subtract(0, 2)).To(Equal(-2))
	})

	It("Should return -2 when 2 is subtracted from 0", func() {
		Expect(Subtract(2, 0)).To(Equal(2))
	})
})

package main

import (
	"math"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestStringCalculator(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "String Calculator Suite")
}

var _ = Describe("Check the inputs", func() {
	It("should return invalid input of input contains alphabets", func() {
		result, valid := CheckInput("aab12")
		Expect(result).To(Equal(math.MaxInt))
		Expect(valid).To(Equal(false))
	})

	It("should return invalid input if input string is of length > 5", func() {
		result, valid := CheckInput("123456")
		Expect(result).To(Equal(math.MaxInt))
		Expect(valid).To(Equal(false))
	})

	It("should return valid input if input is string of length <= 5", func() {
		result, valid := CheckInput("12345")
		Expect(result).To(Equal(12345))
		Expect(valid).To(Equal(true))
	})
})

var _ = Describe("perform the Add Operation", func() {
	It("should return valid output if output is string of length <= 5", func() {
		result, valid := CheckOutput(12345)
		Expect(result).To(Equal("12345"))
		Expect(valid).To(Equal(true))
	})

	It("should return valid output if output is string of length <= 5", func() {
		result, valid := CheckOutput(Add(12345, 12345))
		Expect(result).To(Equal("24690"))
		Expect(valid).To(Equal(true))
	})
})

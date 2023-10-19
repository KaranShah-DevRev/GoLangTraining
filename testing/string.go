package main

import (
	"math"
	"strconv"
)

func CheckInput(input string) (int, bool) {
	if len(input) > 5 {
		return math.MaxInt, false
	}
	inp, err := strconv.Atoi(input)
	if err != nil {
		return math.MaxInt, false
	}
	return inp, true

}

func CheckOutput(output int) (string, bool) {
	op := strconv.Itoa(output)
	return op, true

}

func AddStrings(num1, num2 int) int {
	return num1 + num2
}

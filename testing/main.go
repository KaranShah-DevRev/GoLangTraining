package main

import "math"

func Add(num1 int, num2 int) int {
	data := num1 + num2
	return data
}

func Divide(num1 int, num2 int) int {
	if num2 == 0 {
		return math.MinInt
	}
	data := num1 / num2
	return data
}

func Multiply(num1 int, num2 int) int {
	data := num1 * num2
	return data
}

func Subtract(num1 int, num2 int) int {
	data := num1 - num2
	return data
}

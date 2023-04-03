package controllers

import "errors"

func addition(a, b int) int {
	return a + b
}

func subtraction(a, b int) int {
	return a - b
}

func multiplication(a, b int) int {
	return a * b
}

func division(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

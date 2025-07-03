package utils

import "fmt"

func FindMin(a, b, c float32) float32 {
	min := a
	if b < min {
		min = b
	}
	if c < min {
		min = c
	}
	return min
}

func FindAvg(a, b, c float32) float32 {
	return (a + b + c) / 3
}

func SolveEquation(a, b float32) (float32, error) {
	if a == 0 {
		return 0, fmt.Errorf("рівняння немає рішення (a = 0)")
	}
	return -b / a, nil
}

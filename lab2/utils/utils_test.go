package utils

import (
	"math"
	"testing"
)

func TestFindMin(t *testing.T) {
	got := FindMin(3.5, 7.3, 2.1)
	var want float32 = 2.1
	if got != want {
		t.Errorf("Тест не пройдено! FindMin(3.5, 7.3, 2.1) = %.1f; має бути %.1f", got, want)
	}
}

func TestFindAvg(t *testing.T) {
	got := FindAvg(2.3, 4.6, 7.5)
	var want float32 = 4.8
	const epsilon = 1e-5
	if math.Abs(float64(got-want)) > epsilon {
		t.Errorf("Тест не пройдено! FindAvg(2.3, 4.6, 7.5) = %.1f; має бути %.1f", got, want)
	}
}

func TestSolveEquation(t *testing.T) {
	var a, b float32 = 2, -4
	var want float32 = 2
	const epsilon = 1e-5
	got, _ := SolveEquation(a, b)
	if math.Abs(float64(got-want)) > epsilon {
		t.Errorf("Тест не пройдено! SolveEquation(%.1f, %.1f) = %.4f; має бути %.4f", a, b, got, want)
	}
}

func TestSolveEquationException(t *testing.T) {
	var a, b float32 = 0, 5
	_, err := SolveEquation(a, b)
	if err == nil {
		t.Errorf("Тест не пройдено! SolveEquation(%.1f, %.1f) має викликати помилку", a, b)
	}
}

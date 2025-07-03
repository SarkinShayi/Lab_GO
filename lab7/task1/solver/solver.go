package solver

import (
	"errors"
	"math"
)

type Roots struct {
	X1 float64 `json:"x1,omitempty"`
	X2 float64 `json:"x2,omitempty"`
}

func SolveQuadratic(a, b, c float64) (Roots, error) {
	if a == 0 {
		return Roots{}, errors.New("коефіцієнт a не повинен дорівнювати 0")
	}

	d := b*b - 4*a*c
	switch {
	case d < 0:
		return Roots{}, errors.New("дискримінант менше нуля, рішення немає")
	case d == 0:
		x := -b / (2 * a)
		return Roots{X1: x, X2: x}, nil
	default:
		sqrtD := math.Sqrt(d)
		x1 := (-b + sqrtD) / (2 * a)
		x2 := (-b - sqrtD) / (2 * a)
		return Roots{X1: x1, X2: x2}, nil
	}
}

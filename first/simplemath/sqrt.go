package simplemath

import "math"

func Sqrt(target int) int {
	v := math.Sqrt(float64(target))
	return int(v)
}

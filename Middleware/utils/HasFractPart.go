package utils

import "math"

func HasFractPart(num float32) bool {
	_, frac := math.Modf(float64(num))
	return frac != 0
}

package main

import "math"

func sqrtSum(initial float64, until int) float64 {
	result := initial
	for i := 0; i < until;i++  {
		result += math.Sqrt(float64(i))
	}
	return result
}
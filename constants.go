package main

import (
	"math"
)

var POS_INFINITY float64 = math.Inf(1)

const PI float64 = 3.1415926535897932385

func DegreesToRadians(degrees float64) float64 {
	return degrees * PI / 180.0
}

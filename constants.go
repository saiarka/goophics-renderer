package main

import (
	"errors"
	"math"
	"math/rand"
)

var POS_INFINITY float64 = math.Inf(1)

const PI float64 = 3.1415926535897932385

func DegreesToRadians(degrees float64) float64 {
	return degrees * PI / 180.0
}

func RandomFloat64(bounds ...float64) (float64, error) {
	if len(bounds) == 0 {
		return rand.Float64(), nil
	} else if len(bounds) == 2 {
		return bounds[0] + (bounds[1]-bounds[0])*rand.Float64(), nil
	} else {
		return 0.0, errors.New("Invalid Range for Random Number")
	}
}

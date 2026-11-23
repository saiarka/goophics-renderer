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

func RandomFloat64Interval(i Interval) float64 {
	var f float64

	if i.is_min_inclusive && !i.is_max_inclusive {
		f = i.min + (i.max-i.min)*rand.Float64()
	} else if i.is_min_inclusive && i.is_max_inclusive {
		next_high := math.Nextafter(i.max, POS_INFINITY)
		f = i.min + (next_high-i.min)*rand.Float64()
		if f > i.max {
			f = i.max
		}
	} else if !i.is_min_inclusive && !i.is_max_inclusive {
		nudged_low := math.Nextafter(i.min, i.max)
		f = nudged_low + (i.max-nudged_low)*rand.Float64()
	} else if !i.is_min_inclusive && i.is_max_inclusive {
		f = i.max - (i.max-i.min)*rand.Float64()
	}

	return f
}

func RandomPoint3(x_interval Interval, y_interval Interval, z_interval Interval) Point3 {
	x := RandomFloat64Interval(x_interval)
	y := RandomFloat64Interval(y_interval)
	z := RandomFloat64Interval(z_interval)
	return NewPoint3(x, y, z)
}

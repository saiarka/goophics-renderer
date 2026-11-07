package utils

import (
	"fmt"
)

type Color = Vec3

func NewColor(e1, e2, e3 float64) Vec3 {
	return Vec3{
		e: [3]float64{e1, e2, e3},
	}
}

func WriteColor(color Color) {
	var bounded_r float64 = color.GetVec3X()
	var bounded_g float64 = color.GetVec3Y()
	var bounded_b float64 = color.GetVec3Z()

	r := int(bounded_r * 255.99)
	g := int(bounded_g * 255.99)
	b := int(bounded_b * 255.99)

	fmt.Print(r, " ", g, " ", b, "\n")
}

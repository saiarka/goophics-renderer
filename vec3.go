package main

import (
	"fmt"
	"math"
)

// ///// VECTOR (3) ///////
type Vec3 struct {
	e [3]float64
}

func NewVec3(e1, e2, e3 float64) Vec3 {
	return Vec3{
		e: [3]float64{e1, e2, e3},
	}
}

func (v Vec3) GetVec3X() float64 {
	return v.e[0]
}

func (v Vec3) GetVec3Y() float64 {
	return v.e[1]
}

func (v Vec3) GetVec3Z() float64 {
	return v.e[2]
}

func (v Vec3) GetInverse() Vec3 {
	return Vec3{
		e: [3]float64{-v.e[0], -v.e[1], -v.e[2]},
	}
}

func (v Vec3) Print() {
	fmt.Print(v.e[0], ' ', v.e[1], ' ', v.e[2])
}

// PRIVATE
func (v Vec3) LengthSquared() float64 {
	return v.e[0]*v.e[0] + v.e[1]*v.e[1] + v.e[2]*v.e[2]
}

// PRIVATE
func (v Vec3) length() float64 {
	return math.Sqrt(v.LengthSquared())
}

func (v1 Vec3) AddNewVec3(v2 Vec3) Vec3 {
	return Vec3{
		e: [3]float64{v1.e[0] + v2.e[0], v1.e[1] + v2.e[1], v1.e[2] + v2.e[2]},
	}
}

func (v1 *Vec3) AddVec3(v2 Vec3) {
	v1.e[0] += v2.e[0]
	v1.e[1] += v2.e[1]
	v1.e[2] += v2.e[2]
}

func (v1 Vec3) SubtractNewVec3(v2 Vec3) Vec3 {
	return Vec3{
		e: [3]float64{v1.e[0] - v2.e[0], v1.e[1] - v2.e[1], v1.e[2] - v2.e[2]},
	}
}

func (v1 *Vec3) SubtractVec3(v2 Vec3) {
	v1.e[0] -= v2.e[0]
	v1.e[1] -= v2.e[1]
	v1.e[2] -= v2.e[2]
}

func (v1 Vec3) MultiplyNewVec3(v2 Vec3) Vec3 {
	return Vec3{
		e: [3]float64{
			v1.e[0] * v2.e[0],
			v1.e[1] * v2.e[1],
			v1.e[2] * v2.e[2],
		},
	}
}

func (v1 *Vec3) MultiplyVec3(v2 Vec3) {
	v1.e[0] *= v2.e[0]
	v1.e[1] *= v2.e[1]
	v1.e[2] *= v2.e[2]
}

func (v1 Vec3) MultiplyNewScalar(s float64) Vec3 {
	return Vec3{
		e: [3]float64{
			v1.e[0] * s,
			v1.e[1] * s,
			v1.e[2] * s,
		},
	}
}

func (v1 *Vec3) MultiplyScalar(s float64) {
	v1.e[0] *= s
	v1.e[1] *= s
	v1.e[2] *= s
}

func (v1 Vec3) DivideNewVec3(v2 Vec3) Vec3 {
	return Vec3{
		e: [3]float64{
			v1.e[0] / v2.e[0],
			v1.e[1] / v2.e[1],
			v1.e[2] / v2.e[2],
		},
	}
}

func (v1 *Vec3) DivideVec3(v2 Vec3) {
	v1.e[0] /= v2.e[0]
	v1.e[1] /= v2.e[1]
	v1.e[2] /= v2.e[2]
}

func (v1 Vec3) DivideNewScalar(s float64) Vec3 {
	return Vec3{
		e: [3]float64{
			v1.e[0] / s,
			v1.e[1] / s,
			v1.e[2] / s,
		},
	}
}

func (v1 *Vec3) DivideScalar(s float64) {
	v1.e[0] /= s
	v1.e[1] /= s
	v1.e[2] /= s
}

func (v1 Vec3) DotProduct(v2 Vec3) float64 {
	return v1.e[0]*v2.e[0] + v1.e[1]*v2.e[1] + v1.e[2]*v2.e[2]
}

func (v1 Vec3) CrossProduct(v2 Vec3) Vec3 {
	return Vec3{
		e: [3]float64{
			v1.e[1]*v2.e[2] - v1.e[2]*v2.e[1],
			v1.e[2]*v2.e[0] - v1.e[0]*v2.e[2],
			v1.e[0]*v2.e[1] - v1.e[1]*v2.e[0],
		},
	}
}

func (v Vec3) UnitVector() Vec3 {
	return v.DivideNewScalar(v.length())
}

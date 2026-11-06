package utils

import (
	"fmt"
	"math"
)

// ///// VECTOR CONSTRUCTION WORK ///////
type vec3 struct {
	e [3]float64
}

func NewVec3(e1, e2, e3 float64) vec3 {
	return vec3{
		e: [3]float64{e1, e2, e3},
	}
}

func (v vec3) GetVec3X() float64 {
	return v.e[0]
}

func (v vec3) GetVec3Y() float64 {
	return v.e[1]
}

func (v vec3) GetVec3Z() float64 {
	return v.e[2]
}

func (v vec3) GetInverse() vec3 {
	return vec3{
		e: [3]float64{-v.e[0], -v.e[1], -v.e[2]},
	}
}

func (v vec3) Print() {
	fmt.Print(v.e[0], ' ', v.e[1], ' ', v.e[2])
}

// PRIVATE
func (v vec3) length_squared() float64 {
	return v.e[0]*v.e[0] + v.e[1]*v.e[1] + v.e[2]*v.e[2]
}

// PRIVATE
func (v vec3) length() float64 {
	return math.Sqrt(v.length_squared())
}

func (v1 vec3) AddNewVec3(v2 vec3) vec3 {
	return vec3{
		e: [3]float64{v1.e[0] + v2.e[0], v1.e[1] + v2.e[1], v1.e[2] + v2.e[2]},
	}
}

func (v1 *vec3) AddVec3(v2 vec3) {
	v1.e[0] += v2.e[0]
	v1.e[1] += v2.e[1]
	v1.e[2] += v2.e[2]
}

func (v1 vec3) SubtractNewVec3(v2 vec3) vec3 {
	return vec3{
		e: [3]float64{v1.e[0] - v2.e[0], v1.e[1] - v2.e[1], v1.e[2] - v2.e[2]},
	}
}

func (v1 *vec3) SubtractVec3(v2 vec3) {
	v1.e[0] -= v2.e[0]
	v1.e[1] -= v2.e[1]
	v1.e[2] -= v2.e[2]
}

func (v1 vec3) MultiplyNewVec3(v2 vec3) vec3 {
	return vec3{
		e: [3]float64{
			v1.e[0] * v2.e[0],
			v1.e[1] * v2.e[1],
			v1.e[2] * v2.e[2],
		},
	}
}

func (v1 *vec3) MultiplyVec3(v2 vec3) {
	v1.e[0] *= v2.e[0]
	v1.e[1] *= v2.e[1]
	v1.e[2] *= v2.e[2]
}

func (v1 vec3) MultiplyNewScalar(s float64) vec3 {
	return vec3{
		e: [3]float64{
			v1.e[0] * s,
			v1.e[1] * s,
			v1.e[2] * s,
		},
	}
}

func (v1 *vec3) MultiplyScalar(s float64) {
	v1.e[0] *= s
	v1.e[1] *= s
	v1.e[2] *= s
}

func (v1 vec3) DivideNewVec3(v2 vec3) vec3 {
	return vec3{
		e: [3]float64{
			v1.e[0] / v2.e[0],
			v1.e[1] / v2.e[1],
			v1.e[2] / v2.e[2],
		},
	}
}

func (v1 *vec3) DivideVec3(v2 vec3) {
	v1.e[0] /= v2.e[0]
	v1.e[1] /= v2.e[1]
	v1.e[2] /= v2.e[2]
}

func (v1 vec3) DivideNewScalar(s float64) vec3 {
	return vec3{
		e: [3]float64{
			v1.e[0] / s,
			v1.e[1] / s,
			v1.e[2] / s,
		},
	}
}

func (v1 *vec3) DivideScalar(s float64) {
	v1.e[0] /= s
	v1.e[1] /= s
	v1.e[2] /= s
}

func (v1 vec3) DotProduct(v2 vec3) float64 {
	return v1.e[0]*v2.e[0] + v1.e[1]*v2.e[1] + v1.e[2]*v2.e[2]
}

func (v1 vec3) CrossProduct(v2 vec3) vec3 {
	return vec3{
		e: [3]float64{
			v1.e[1]*v2.e[2] - v1.e[2]*v2.e[1],
			v1.e[2]*v2.e[0] - v1.e[0]*v2.e[2],
			v1.e[0]*v2.e[1] - v1.e[1]*v2.e[0],
		},
	}
}

func (v vec3) UnitVector() vec3 {
	return v.DivideNewScalar(v.length())
}

/////// VECTOR CONSTRUCTION WORK ///////

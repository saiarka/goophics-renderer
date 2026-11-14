package main

import (
	"math"
)

// SPHERE//

type Sphere struct {
	location Point3
	radius   float64
}

func (s Sphere) CheckHit(r Ray, t_min float64, t_max float64, hit_record *HitRecord) bool {
	center_to_point := s.location.SubtractNewVec3(r.Origin())
	a := r.Direction().LengthSquared()
	h := r.Direction().DotProduct(center_to_point)
	c := center_to_point.LengthSquared() - (s.radius * s.radius)
	discriminant := (h * h) - (a * c)

	if discriminant < 0 {
		return NOT_HIT
	}

	t := (h - math.Sqrt(discriminant)) / a
	if t <= t_min || t >= t_max {
		t = (h + math.Sqrt(discriminant)) / a
		if t <= t_min || t >= t_max {
			return NOT_HIT
		}
	}

	hit_record.point = r.At(t)
	hit_record.normal = hit_record.point.SubtractNewVec3(s.location).UnitVector()
	hit_record.t = t

	return IS_HIT
}

//SPHERE//

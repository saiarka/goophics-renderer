package main

// HITTABLE INTERFACE //

const (
	IS_HIT  = true
	NOT_HIT = false
)

type HitRecord struct {
	point  Point3
	normal Vec3
	t      float64
}

type Hittable interface {
	CheckHit(r Ray, t_min float64, t_max float64, hit_record *HitRecord) bool
}

// HITTABLE INTERFACE //

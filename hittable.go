package main

// HITTABLE INTERFACE //

const (
	IS_HIT  = true
	NOT_HIT = false
)

type HitRecord struct {
	point         Point3
	normal        Vec3
	t             float64
	is_front_face bool
}

type Hittable interface {
	CheckHit(r Ray, i Interval, hit_record *HitRecord) (is_hit bool)
}

// HITTABLE INTERFACE //

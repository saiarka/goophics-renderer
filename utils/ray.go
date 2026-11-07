package utils

type Point3 = Vec3

func NewPoint3(e1, e2, e3 float64) Vec3 {
	return Vec3{
		e: [3]float64{e1, e2, e3},
	}
}

type Ray struct {
	origin    Point3
	direction Vec3
}

func NewRay(origin Point3, direction Vec3) Ray {
	return Ray{
		origin:    NewVec3(origin.e[0], origin.e[1], origin.e[2]),
		direction: NewVec3(direction.e[0], direction.e[1], direction.e[2]),
	}
}

func (r Ray) Origin() Point3 {
	return r.origin
}

func (r Ray) Direction() Vec3 {
	return r.direction
}

func (r Ray) At(t float64) Point3 {
	return r.origin.AddNewVec3(r.direction.MultiplyNewScalar(t))
}

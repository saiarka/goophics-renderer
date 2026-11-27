// Describes different rendering configurations with respective settings
// To be passed into the camera functions

package main

type Renderer struct {
	aa_config      AntialiasConfig
	diffuse_config DiffuseConfig
}

type AntialiasConfig struct {
	sample_size int
	enabled     bool
}

type DiffuseConfig struct {
	attenuation  float64
	distribution string //TODO: maybe?
	enabled      bool
}

func rayColor(r Ray, world HittableList) Color {
	h_rec := HitRecord{
		nil,
		NewPoint3(0, 0, 0),
		NewVec3(0, 0, 0),
		0,
		false,
	}

	if world.CheckHit(r, NewInterval(0, POS_INFINITY, false, false), &h_rec) {
		return h_rec.normal.AddNewVec3(NewColor(1, 1, 1)).MultiplyNewScalar(0.5)
	}

	unit_direction := r.Direction().UnitVector()
	a := (unit_direction.GetVec3Y() + 1.0) * 0.5
	color_1_portion := NewColor(1.0, 1.0, 1.0).MultiplyNewScalar(1.0 - a)
	color_2_portion := NewColor(0.5, 0.7, 1.0).MultiplyNewScalar(a)
	return color_1_portion.AddNewVec3(color_2_portion)
}

func rayColorDiffuse(r Ray, world HittableList, attenuation float64) Color {
	h_rec := HitRecord{
		nil,
		NewPoint3(0, 0, 0),
		NewVec3(0, 0, 0),
		0,
		false,
	}

	if world.CheckHit(r, NewInterval(0, POS_INFINITY, false, false), &h_rec) {
		switch (*h_rec.object_ptr).(type) {
		case Sphere:
			dir := GenerateRandomSpherePoint(
				Interval{
					-1,
					1,
					true,
					true,
				},
				Interval{
					-1,
					1,
					true,
					true,
				},
				Interval{
					-1,
					1,
					true,
					true,
				},
			)
			if hemisphere_check := dir.DotProduct(h_rec.normal); hemisphere_check <= 0 {
				dir = dir.GetInverse()
			}
			return rayColorDiffuse(NewRay(h_rec.point, dir), world, attenuation).MultiplyNewScalar(attenuation)
		}
	}
	unit_direction := r.Direction().UnitVector()
	a := (unit_direction.GetVec3Y() + 1.0) * 0.5
	color_1_portion := NewColor(1.0, 1.0, 1.0).MultiplyNewScalar(1.0 - a)
	color_2_portion := NewColor(0.5, 0.7, 1.0).MultiplyNewScalar(a)
	return color_1_portion.AddNewVec3(color_2_portion)
}

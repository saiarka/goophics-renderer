package main

import (
	"fmt"
	"os"
)

var IMAGE_WIDTH int = 800
var ASPECT_RATIO float64 = 16.0 / 9.0

func RayColor(r Ray) Color {
	unit_direction := r.Direction().UnitVector()
	a := (unit_direction.GetVec3Y() + 1.0) * 0.5
	color_1_portion := NewColor(1.0, 1.0, 1.0).MultiplyNewScalar(1.0 - a)
	color_2_portion := NewColor(0.5, 0.7, 1.0).MultiplyNewScalar(a)
	return color_1_portion.AddNewVec3(color_2_portion)
}

var SPHERE_COLOR Color = NewColor(255, 0, 0)
var SPHERE_LOCATION Point3 = NewPoint3(0, 0, -1)

func main() {

	IMAGE_HEIGHT := int(float64(IMAGE_WIDTH) / ASPECT_RATIO)
	IMAGE_WIDTH = max(1, IMAGE_WIDTH)

	//CAMERA
	// NOTE: 'Camera' really refers to a set of vectors and dimension standards referring to the
	// actual rendering screen and coordinate system used to understand pixel location
	focal_length := 1.0
	var VIEWPORT_HEIGHT float64 = 2.0
	VIEWPORT_WIDTH := VIEWPORT_HEIGHT * (float64(IMAGE_WIDTH) / float64(IMAGE_HEIGHT))
	var camera_center Point3 = NewPoint3(0, 0, 0)

	var viewport_u Vec3 = NewVec3(VIEWPORT_WIDTH, 0, 0)
	var viewport_v Vec3 = NewVec3(0, -VIEWPORT_HEIGHT, 0) // Our coodinate system has negative y going down

	var pixel_delta_u Vec3 = viewport_u.DivideNewScalar(float64(IMAGE_WIDTH))  // viewport width divided by # of pixels width-wise
	var pixel_delta_v Vec3 = viewport_v.DivideNewScalar(float64(IMAGE_HEIGHT)) // viewport height divided by # of pixels height-wise

	var viewport_upper_left Point3 = camera_center.SubtractNewVec3(NewVec3(0, 0, focal_length)).SubtractNewVec3(viewport_u.DivideNewScalar(2)).SubtractNewVec3(viewport_v.DivideNewScalar(2))
	var pixel_00_loc Point3 = pixel_delta_u.AddNewVec3(pixel_delta_v).MultiplyNewScalar(0.5).AddNewVec3(viewport_upper_left) //Calculating location of 'first' pixel (top left of viewport)

	fmt.Print("P3\n")
	fmt.Print(IMAGE_WIDTH, " ", IMAGE_HEIGHT, "\n")
	fmt.Print("255\n")

	s := Sphere{
		SPHERE_LOCATION,
		0.5,
	}

	for i := range IMAGE_HEIGHT {
		fmt.Fprint(os.Stderr, "Scanlines Remaining: ", IMAGE_HEIGHT-i, "\n")
		for j := range IMAGE_WIDTH {
			var pixel_center Point3 = pixel_00_loc.AddNewVec3(pixel_delta_u.MultiplyNewScalar(float64(j))).AddNewVec3(pixel_delta_v.MultiplyNewScalar(float64(i)))
			ray_direction := pixel_center.SubtractNewVec3(camera_center)
			var r Ray = NewRay(camera_center, ray_direction)

			h_rec := HitRecord{
				NewPoint3(0, 0, 0),
				NewVec3(0, 0, 0),
				0,
			}
			if check_hit := s.CheckHit(r, -1000, 1000, &h_rec); check_hit == IS_HIT {
				WriteColor(NewColor(h_rec.normal.GetVec3X()+1, h_rec.normal.GetVec3Y()+1, h_rec.normal.GetVec3Z()+1).MultiplyNewScalar(0.5))
			} else {
				var pixel_color Color = RayColor(r)
				WriteColor(pixel_color)
			}
		}
	}

}

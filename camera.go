package main

import (
	"fmt"
	"os"
)

type Camera struct {
	focal_length    float64
	viewport_height float64
	viewport_width  float64
	camera_center   Point3
}

type CameraOption func(c *Camera)

func WithFocalLength(f float64) CameraOption {
	return func(c *Camera) {
		c.focal_length = f
	}
}

func WithViewportHeight(h float64) CameraOption {
	return func(c *Camera) {
		c.viewport_height = h
	}
}

func WithCameraCenter(p Point3) CameraOption {
	return func(c *Camera) {
		c.camera_center = p
	}
}

func NewCamera(camera_opts ...CameraOption) Camera {
	//Default Camera Values
	c := Camera{
		focal_length:    1.0,
		viewport_height: 2.0,
		viewport_width:  2.0 * (float64(IMAGE_WIDTH) / float64(IMAGE_HEIGHT)),
		camera_center:   NewPoint3(0, 0, 0),
	}

	for _, opts := range camera_opts {
		opts(&c)
	}

	c.viewport_width = c.viewport_height * (float64(IMAGE_WIDTH) / float64(IMAGE_HEIGHT))

	return c
}

// TODO: We should consider moving ViewportConfig related code to a seperate file
type ViewportConfig struct {
	pixel_00_loc  Point3
	pixel_delta_u Vec3
	pixel_delta_v Vec3
}

func (c *Camera) getViewportConfig() ViewportConfig {

	var viewport_u Vec3 = NewVec3(c.viewport_width, 0, 0)
	var viewport_v Vec3 = NewVec3(0, -c.viewport_height, 0) // Our coodinate system has negative y going down

	var pixel_delta_u Vec3 = viewport_u.DivideNewScalar(float64(IMAGE_WIDTH))  // viewport width divided by # of pixels width-wise
	var pixel_delta_v Vec3 = viewport_v.DivideNewScalar(float64(IMAGE_HEIGHT)) // viewport height divided by # of pixels height-wise

	var viewport_upper_left Point3 = c.camera_center.SubtractNewVec3(NewVec3(0, 0, c.focal_length)).SubtractNewVec3(viewport_u.DivideNewScalar(2)).SubtractNewVec3(viewport_v.DivideNewScalar(2))
	var pixel_00_loc Point3 = pixel_delta_u.AddNewVec3(pixel_delta_v).MultiplyNewScalar(0.5).AddNewVec3(viewport_upper_left) //Calculating location of 'first' pixel (top left of viewport)

	return ViewportConfig{
		pixel_00_loc:  pixel_00_loc,
		pixel_delta_u: pixel_delta_u,
		pixel_delta_v: pixel_delta_v,
	}

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

func rayColorDiffuse(r Ray, world HittableList, attenuation float64) Color { //TODO: Replace with RenderConfiguration
	h_rec := HitRecord{
		nil,
		NewPoint3(0, 0, 0),
		NewVec3(0, 0, 0),
		0,
		false,
	}

	if world.CheckHit(r, NewInterval(0, POS_INFINITY, false, false), &h_rec) {
		// TODO: Below is a coloring method using normals
		// We want to specify rendering styles depending on material of the object eventually
		// return h_rec.normal.AddNewVec3(NewColor(1, 1, 1)).MultiplyNewScalar(0.5)

		switch (*h_rec.object_ptr).(type) {
		case Sphere:
			//Generates random point3 in sphere
			//Check if length of vector squared overflows
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
			//Creates vector from sphere origin to point3
			// dir := random_point.SubtractNewVec3(obj.location)
			//Checks if in correct hemisphere by comparing to normal --> If not, then reverse
			if hemisphere_check := dir.DotProduct(h_rec.normal); hemisphere_check <= 0 {
				dir = dir.GetInverse()
			}
			//Recursive check on new random ray, with respective altered attenuation factor
			return rayColorDiffuse(NewRay(h_rec.point, dir), world, attenuation).MultiplyNewScalar(attenuation)
			// return h_rec.normal.AddNewVec3(NewColor(1, 1, 1)).MultiplyNewScalar(0.5)
		}

	}

	unit_direction := r.Direction().UnitVector()
	a := (unit_direction.GetVec3Y() + 1.0) * 0.5
	color_1_portion := NewColor(1.0, 1.0, 1.0).MultiplyNewScalar(1.0 - a)
	color_2_portion := NewColor(0.5, 0.7, 1.0).MultiplyNewScalar(a)
	return color_1_portion.AddNewVec3(color_2_portion)
}

func (c *Camera) Render(world HittableList) {
	vc := c.getViewportConfig()

	for i := range IMAGE_HEIGHT {
		fmt.Fprint(os.Stderr, "Scanlines Remaining: ", IMAGE_HEIGHT-i, "\n")
		for j := range IMAGE_WIDTH {
			var pixel_center Point3 = vc.pixel_00_loc.AddNewVec3(vc.pixel_delta_u.MultiplyNewScalar(float64(j))).AddNewVec3(vc.pixel_delta_v.MultiplyNewScalar(float64(i)))
			ray_direction := pixel_center.SubtractNewVec3(c.camera_center)
			var r Ray = NewRay(c.camera_center, ray_direction)

			var pixel_color Color = rayColor(r, world)
			WriteColor(pixel_color)

		}
	}

}

func (c *Camera) RenderDiffuse(world HittableList) {
	vc := c.getViewportConfig()

	for i := range IMAGE_HEIGHT {
		fmt.Fprint(os.Stderr, "Scanlines Remaining: ", IMAGE_HEIGHT-i, "\n")
		for j := range IMAGE_WIDTH {
			var pixel_center Point3 = vc.pixel_00_loc.AddNewVec3(vc.pixel_delta_u.MultiplyNewScalar(float64(j))).AddNewVec3(vc.pixel_delta_v.MultiplyNewScalar(float64(i)))
			ray_direction := pixel_center.SubtractNewVec3(c.camera_center)
			var r Ray = NewRay(c.camera_center, ray_direction)

			var pixel_color Color = rayColorDiffuse(r, world, 0.5)
			WriteColor(pixel_color)

		}
	}

}

func (c *Camera) RenderWithAA(world HittableList, samples int) { //TODO: Remove and replace with RenderConfiguration
	vc := c.getViewportConfig()

	for i := range IMAGE_HEIGHT {
		fmt.Fprint(os.Stderr, "Scanlines Remaining: ", IMAGE_HEIGHT-i, "\n")
		for j := range IMAGE_WIDTH {
			var pixel_center Point3 = vc.pixel_00_loc.AddNewVec3(vc.pixel_delta_u.MultiplyNewScalar(float64(j))).AddNewVec3(vc.pixel_delta_v.MultiplyNewScalar(float64(i)))
			var accumulation Color = NewColor(0, 0, 0)
			for range samples {
				random_x, err := RandomFloat64(-vc.pixel_delta_u.length()/2, vc.pixel_delta_u.length()/2)
				if err != nil {
					fmt.Print(err)
				}
				random_y, err := RandomFloat64(-vc.pixel_delta_v.length()/2, vc.pixel_delta_v.length()/2)
				if err != nil {
					fmt.Print(err)
				}
				var sample_point Point3 = pixel_center.AddNewVec3(NewVec3(random_x, random_y, 0))
				ray_direction := sample_point.SubtractNewVec3(c.camera_center)
				var r Ray = NewRay(c.camera_center, ray_direction)
				accumulation.AddVec3(rayColor(r, world))
			}
			accumulation.DivideScalar(float64(samples))
			WriteColor(accumulation)
		}
	}
}

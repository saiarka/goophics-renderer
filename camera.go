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

func (c *Camera) Render(world HittableList, renderer *Renderer) {
	vc := c.getViewportConfig()

	if renderer.aa_config.enabled {
		for i := range IMAGE_HEIGHT {
			fmt.Fprint(os.Stderr, "Scanlines Remaining: ", IMAGE_HEIGHT-i, "\n")
			for j := range IMAGE_WIDTH {
				var pixel_center Point3 = vc.pixel_00_loc.AddNewVec3(vc.pixel_delta_u.MultiplyNewScalar(float64(j))).AddNewVec3(vc.pixel_delta_v.MultiplyNewScalar(float64(i)))
				var accumulation Color = NewColor(0, 0, 0)
				for range renderer.aa_config.sample_size {
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
					if renderer.diffuse_config.enabled {
						accumulation.AddVec3(rayColorDiffuse(r, world, renderer.diffuse_config.attenuation))
					} else {
						accumulation.AddVec3(rayColor(r, world))
					}
				}
				accumulation.DivideScalar(float64(renderer.aa_config.sample_size))
				WriteColor(accumulation)
			}
		}
	} else {
		for i := range IMAGE_HEIGHT {
			fmt.Fprint(os.Stderr, "Scanlines Remaining: ", IMAGE_HEIGHT-i, "\n")
			for j := range IMAGE_WIDTH {
				var pixel_center Point3 = vc.pixel_00_loc.AddNewVec3(vc.pixel_delta_u.MultiplyNewScalar(float64(j))).AddNewVec3(vc.pixel_delta_v.MultiplyNewScalar(float64(i)))
				ray_direction := pixel_center.SubtractNewVec3(c.camera_center)
				var r Ray = NewRay(c.camera_center, ray_direction)

				var pixel_color Color
				if renderer.diffuse_config.enabled {
					pixel_color = rayColorDiffuse(r, world, renderer.diffuse_config.attenuation)
				} else {
					pixel_color = rayColor(r, world)
				}
				WriteColor(pixel_color)

			}
		}
	}

}

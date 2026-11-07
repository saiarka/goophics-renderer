package main

import (
	"fmt"
	utils "github.com/saiarka/goophics-renderer/utils"
	"os"
)

var IMAGE_WIDTH int = 400
var ASPECT_RATIO float64 = 16.0 / 9.0

func RayColor(r utils.Ray) utils.Color {
	return utils.NewColor(210, 0, 0)
}

func main() {

	IMAGE_HEIGHT := int(float64(IMAGE_WIDTH) / ASPECT_RATIO)
	IMAGE_WIDTH = max(1, IMAGE_WIDTH)

	//CAMERA
	// NOTE: 'Camera' really refers to a set of vectors and dimension standards referring to the
	// actual rendering screen and coordinate system used to understand pixel location
	focal_length := 1.0
	var VIEWPORT_HEIGHT float64 = 2.0
	VIEWPORT_WIDTH := VIEWPORT_HEIGHT * (float64(IMAGE_WIDTH) / float64(IMAGE_HEIGHT))
	var camera_center utils.Point3 = utils.NewPoint3(0, 0, 0)

	var viewport_u utils.Vec3 = utils.NewVec3(VIEWPORT_WIDTH, 0, 0)
	var viewport_v utils.Vec3 = utils.NewVec3(0, -VIEWPORT_HEIGHT, 0) // Our coodinate system has negative y going down

	var pixel_delta_u utils.Vec3 = viewport_u.DivideNewScalar(float64(IMAGE_WIDTH))  // viewport width divided by # of pixels width-wise
	var pixel_delta_v utils.Vec3 = viewport_v.DivideNewScalar(float64(IMAGE_HEIGHT)) // viewport height divided by # of pixels height-wise

	var viewport_upper_left utils.Point3 = camera_center.SubtractNewVec3(utils.NewVec3(0, 0, focal_length)).SubtractNewVec3(viewport_u.DivideNewScalar(2)).SubtractNewVec3(viewport_v.DivideNewScalar(2))
	var pixel_00_loc utils.Point3 = pixel_delta_u.AddNewVec3(pixel_delta_v).MultiplyNewScalar(0.5).AddNewVec3(viewport_upper_left) //Calculating location of 'first' pixel (top left of viewport)

	fmt.Print("P3\n")
	fmt.Print(IMAGE_WIDTH, " ", IMAGE_HEIGHT, "\n")
	fmt.Print("255\n")

	for i := range IMAGE_HEIGHT {
		fmt.Fprint(os.Stderr, "Scanlines Remaining: ", IMAGE_HEIGHT-i, "\n")
		for j := range IMAGE_WIDTH {
			var pixel_center utils.Point3 = pixel_00_loc.AddNewVec3(pixel_delta_u.MultiplyNewScalar(float64(j))).AddNewVec3(pixel_delta_v.MultiplyNewScalar(float64(i)))
			ray_direction := pixel_center.SubtractNewVec3(camera_center)
			var r utils.Ray = utils.NewRay(camera_center, ray_direction)

			var pixel_color utils.Color = RayColor(r)
			utils.WriteColor(pixel_color)
		}
	}

}

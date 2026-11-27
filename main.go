package main

import (
	"flag"
	"fmt"
)

var IMAGE_WIDTH int = 800
var IMAGE_HEIGHT int = int(float64(IMAGE_WIDTH) / ASPECT_RATIO)
var ASPECT_RATIO float64 = 16.0 / 9.0

var SPHERE_COLOR Color = NewColor(255, 0, 0)
var SPHERE_LOCATION Point3 = NewPoint3(0, 0, -1)

func main() {

	//NOTE: Leaving this here bc I forgot why I had it in the first place
	// IMAGE_WIDTH = max(1, IMAGE_WIDTH)

	// Flags for rendering configuration
	// SUPPORTS:
	// AA (Antialiasing)
	antialiasing_enabled := flag.Bool("aa", false, "bool flag to enable antialiasing")

	flag.Parse()

	camera := NewCamera()

	fmt.Print("P3\n")
	fmt.Print(IMAGE_WIDTH, " ", IMAGE_HEIGHT, "\n")
	fmt.Print("255\n")

	world := HittableList{}
	world.Add(Sphere{
		SPHERE_LOCATION,
		0.5,
	})
	world.Add(Sphere{
		NewPoint3(0, -100.5, -1),
		100,
	})

	if *antialiasing_enabled {
		camera.RenderWithAA(world, 4)
	} else {
		camera.RenderDiffuse(world)
	}

}

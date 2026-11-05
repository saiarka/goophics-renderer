package main

import (
	"fmt"
)

var IMAGE_HEIGHT int = 400
var IMAGE_WIDTH int = 400

func main() {

	fmt.Print("P3\n")
	fmt.Print(IMAGE_WIDTH, " ", IMAGE_HEIGHT, "\n")
	fmt.Print("255\n")

	for i := 0; i < IMAGE_HEIGHT; i++ {
		for j := 0; j < IMAGE_WIDTH; j++ {
			var bounded_r float64 = float64(i) / float64(IMAGE_HEIGHT)
			var bounded_g float64 = float64(j) / float64(IMAGE_WIDTH)
			var bounded_b float64 = 0

			r := int(bounded_r * 255.99)
			g := int(bounded_g * 255.99)
			b := int(bounded_b * 255.99)

			fmt.Print(r, " ", g, " ", b, "\n")
		}
	}

}

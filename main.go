package main

import (
	"fmt"
	utils "github.com/saiarka/goophics-renderer/utils"
	"os"
)

var IMAGE_HEIGHT int = 400
var IMAGE_WIDTH int = 400

func main() {

	fmt.Print("P3\n")
	fmt.Print(IMAGE_WIDTH, " ", IMAGE_HEIGHT, "\n")
	fmt.Print("255\n")

	for i := 0; i < IMAGE_HEIGHT; i++ {
		fmt.Fprint(os.Stderr, "Scanlines Remaining: ", IMAGE_HEIGHT-i, "\n")
		for j := 0; j < IMAGE_WIDTH; j++ {
			pixel_color := utils.NewColor(float64(i)/float64(IMAGE_HEIGHT), float64(j)/float64(IMAGE_WIDTH), 0)
			utils.WriteColor(pixel_color)
		}
	}

}

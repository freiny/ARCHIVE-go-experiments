package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"math"
	"math/rand"
)

func main() {
	rgba := getRGBARandom(4, 4)
	pix := rgba.Pix
	for i := 0; i < len(pix); i += 4 {
		if math.Mod(float64(i), 16) == 0 {
			fmt.Println()
		}
		fmt.Println("RGBA: ", pix[i+0], pix[i+1], pix[i+2], pix[i+3])
	}
}

func getRGBARandom(width int, height int) *image.RGBA {

	// size := size
	rgba := image.NewRGBA(image.Rect(0, 0, width, height))
	blue := color.RGBA{0, 0, 255, 255}
	draw.Draw(rgba, rgba.Bounds(), &image.Uniform{blue}, image.ZP, draw.Src)

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			r := uint8(rand.Intn(255))
			g := uint8(rand.Intn(255))
			b := uint8(rand.Intn(255))
			rgba.Set(x, y, color.RGBA{r, g, b, 255})
		}
	}

	return rgba
}

package mandelbrot

import (
	"image/color"
)

func fillPixel(pixel Pixel) {
	imgSerial.Set(pixel.x, pixel.y, color.RGBA{R: pixel.cr, G: pixel.cg, B: pixel.cb, A: 255})
}

func calculateFractal(args ...float64) {
	x0 := args[0]
	y0 := args[1]
	x1 := args[2]
	y1 := args[3]

	xStep := (x1 - x0) / width
	yStep := (y1 - y0) / height

	for j := 0.0; j < height; j++ {
		for i := 0.0; i < width; i++ {
			colorPoint := getColorPoint(x0+i*xStep, y1-j*yStep, maxIter)
			newPixel := Pixel{x: int(i), y: int(j), cr: colorPoint.R, cg: colorPoint.G, cb: colorPoint.B}
			fillPixel(newPixel)
		}
	}
}

package mandelbrot

import (
	"image"
	"image/color"
)

func fillPixel(pixel Pixel, img *image.RGBA) {
	img.SetRGBA(int(pixel.x), int(pixel.y), color.RGBA{R: pixel.cr, G: pixel.cg, B: pixel.cb, A: 255})
}

func calculateFractal() {
	x0, y0, x1, y1 := args[0], args[1], args[2], args[3]

	xStep := (x1 - x0) / width
	yStep := (y1 - y0) / height

	for j := 0.0; j < height; j++ {
		for i := 0.0; i < width; i++ {
			colorPoint := getColorPoint(x0+i*xStep, y1-j*yStep, maxIter)
			newPixel := Pixel{x: i, y: j, cr: colorPoint.R, cg: colorPoint.G, cb: colorPoint.B}
			fillPixel(newPixel, imgSerial)
		}
	}
}

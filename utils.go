package mandelbrot

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func generatePNG(name string, img *image.RGBA) error {
	file, err := os.Create(fmt.Sprintf("%vx%v-%s.png", width, height, name))
	if err != nil {
		return err
	}

	err = png.Encode(file, img)
	if err != nil {
		return err
	}
	err = file.Close()
	if err != nil {
		return err
	}

	return err
}

func getColorPoint(x, y, maxIter float64) color.RGBA {
	i := 0
	z := complex(0.0, 0.0)

	for i < int(maxIter) {
		z = z*z + complex(x, y)
		if cmplx.Abs(z) > 2.0 {
			break
		}
		i = i + 1
	}
	return color.RGBA{R: uint8(i), G: uint8(i), B: uint8(i), A: 255}
}

package mandelbrot

import "image"

const (
	width      = 1024.0
	height     = 1024.0
	totalPixel = width * height
	maxIter    = 1000
	numBlocks  = 64
	numThreads = 16
)

var (
	args        = []float64{-2.0, -1.0, 1.0, 1.0}
	imgParallel = image.NewRGBA(image.Rect(0, 0, width, height))
	imgSerial   = image.NewRGBA(image.Rect(0, 0, width, height))
	pixelCount  = 0
)

type Pixel struct {
	x  int
	y  int
	cr uint8
	cg uint8
	cb uint8
}

type WorkItem struct {
	initialX int
	finalX   int
	initialY int
	finalY   int
}

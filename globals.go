package mandelbrot

import "image"

// maxIter := 1000
// samples := 200
// numBlocks := 64
// numThreads := 16

const (
	width      = 1024.0
	height     = 1024.0
	totalPixel = width * height
	maxIter    = 100000
	numBlocks  = 64
	numThreads = 16
)

var (
	args        = []float64{-2.0, -1.0, 1.0, 1.0}
	imgParallel = image.NewRGBA(image.Rect(0, 0, width, height))
	imgSerial   = image.NewRGBA(image.Rect(0, 0, width, height))
)

type Pixel struct {
	x  float64
	y  float64
	cr uint8
	cg uint8
	cb uint8
}

type WorkItem struct {
	initialX float64
	finalX   float64
	initialY float64
	finalY   float64
}

package mandelbrot

import (
	"math"
)

func workBufferInit(workBuffer chan WorkItem) {
	var sqrt = math.Sqrt(float64(numBlocks))

	for i := sqrt - 1.0; i >= 0.0; i-- {
		for j := 0.0; j < sqrt; j++ {
			workBuffer <- WorkItem{
				initialX: i * (width / sqrt),
				finalX:   (i + 1) * (width / sqrt),
				initialY: j * (height / sqrt),
				finalY:   (j + 1) * (height / sqrt),
			}
		}
	}
}

func drawThread(drawBuffer chan Pixel) {
	for i := range drawBuffer {
		fillPixel(i, imgParallel)
	}
}

func workersInit(drawBuffer chan Pixel, workBuffer chan WorkItem, threadBuffer chan bool) {
	for i := 1; i <= numThreads; i++ {
		threadBuffer <- true
	}

	for range threadBuffer {
		workItem := <-workBuffer
		go workerThread(workItem, drawBuffer, threadBuffer)
	}
}

func workerThread(workItem WorkItem, drawBuffer chan Pixel, threadBuffer chan bool) {
	x0, y0, x1, y1 := args[0], args[1], args[2], args[3]

	xStep := (x1 - x0) / width
	yStep := (y1 - y0) / height

	for j := workItem.initialY; j < workItem.finalY; j++ {
		for i := workItem.initialX; i < workItem.finalX; i++ {
			colorPoint := getColorPoint(x0+i*xStep, y1-j*yStep, maxIter)
			drawBuffer <- Pixel{x: i, y: j, cr: colorPoint.R, cg: colorPoint.G, cb: colorPoint.B}
		}
	}
	threadBuffer <- true
}

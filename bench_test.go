package mandelbrot

import (
	"fmt"
	"testing"
)

func BenchmarkMandelbrot(b *testing.B) {
	b.Run(fmt.Sprintf("%vx%v-%s", width, height, "sequential"), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			calculateFractal()
			err := generatePNG("sequential", imgSerial)
			if err != nil {
				b.Fatalf(`CreatePNG(fractal) SEQ= %v, error`, err)
			}
		}
	})

	b.Run(fmt.Sprintf("%vx%v-%s", width, height, "parallel"), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			workBuffer := make(chan WorkItem, numBlocks)
			threadBuffer := make(chan bool, numThreads)
			drawBuffer := make(chan Pixel, totalPixel)

			workBufferInit(workBuffer)
			go workersInit(drawBuffer, workBuffer, threadBuffer)
			go drawThread(drawBuffer)

		}
		err := generatePNG("parallel", imgParallel)
		if err != nil {
			b.Fatalf(`generatePNG(img) %e`, err)
		}
	})
}

# Mandelbrot Set - Sequential vs Parallel
The project has the following structure:

- **go.mod**: Golang Package Manager
- **parallel-mandelbrot.go**: Set of functions to calculate the Mandelbrot Set parallel.
- **serial-mandelbrot.go**: Set of functions to calculate the Mandelbrot set sequentially.
- **globals.go**: Set of global structs and variables.
- **utils.go**: Set of functions used by both programs.
- **bench_test.go**: Golang Benchmark where sequential and parallel programs are executed.

package hw7

import (
	"image"
	"image/color"
	"math/cmplx"
	"runtime"
	"sync"
)

var workerCount = runtime.NumCPU()

// Render computes a Mandelbrot image using either pixel or row based jobs.
// It returns the generated image for benchmarking purposes.
func Render(usePixelJobs bool) *image.RGBA {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	if usePixelJobs {
		renderPixels(img, xmin, ymin, xmax, ymax, width, height)
	} else {
		renderRows(img, xmin, ymin, xmax, ymax, width, height)
	}

	return img
}

// RenderSequential performs the Mandelbrot computation using a single
// goroutine. It mirrors the original implementation for baseline timings.
func RenderSequential() *image.RGBA {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			img.Set(px, py, mandelbrot(complex(x, y)))
		}
	}
	return img
}

func renderRows(img *image.RGBA, xmin, ymin, xmax, ymax float64, width, height int) {
	w := float64(width)
	h := float64(height)
	rows := make(chan int, height)
	var wg sync.WaitGroup
	wg.Add(workerCount)
	for i := 0; i < workerCount; i++ {
		go func() {
			defer wg.Done()
			for py := range rows {
				y := float64(py)/h*(ymax-ymin) + ymin
				for px := 0; px < width; px++ {
					x := float64(px)/w*(xmax-xmin) + xmin
					img.Set(px, py, mandelbrot(complex(x, y)))
				}
			}
		}()
	}
	for py := 0; py < height; py++ {
		rows <- py
	}
	close(rows)
	wg.Wait()
}

func renderPixels(img *image.RGBA, xmin, ymin, xmax, ymax float64, width, height int) {
	type pixel struct{ x, y int }
	w := float64(width)
	h := float64(height)
	jobs := make(chan pixel, 1024)
	var wg sync.WaitGroup
	wg.Add(workerCount)
	for i := 0; i < workerCount; i++ {
		go func() {
			defer wg.Done()
			for p := range jobs {
				y := float64(p.y)/h*(ymax-ymin) + ymin
				x := float64(p.x)/w*(xmax-xmin) + xmin
				img.Set(p.x, p.y, mandelbrot(complex(x, y)))
			}
		}()
	}
	for py := 0; py < height; py++ {
		for px := 0; px < width; px++ {
			jobs <- pixel{px, py}
		}
	}
	close(jobs)
	wg.Wait()
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

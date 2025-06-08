package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
	"runtime"
	"sync"
)

var workerCount = runtime.NumCPU()

const usePixelJobs = false

func main() {
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

	file, err := os.Create("mandelbrot.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	if err := png.Encode(file, img); err != nil {
		panic(err)
	}
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

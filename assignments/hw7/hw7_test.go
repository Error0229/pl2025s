package hw7

import "testing"

func BenchmarkRows(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Render(false)
	}
}

func BenchmarkSequential(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RenderSequential()
	}
}

func BenchmarkPixels(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Render(true)
	}
}

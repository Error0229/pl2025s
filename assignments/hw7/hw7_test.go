package hw7

import "testing"

func BenchmarkRows1worker(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Render(false, 1)
	}
}

func BenchmarkRows4workers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Render(false, 4)
	}
}
func BenchmarkRows8workers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Render(false, 8)
	}
}

func BenchmarkRows12workers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Render(false, 12)
	}
}
func BenchmarkRows16workers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Render(false, 16)
	}
}
func BenchmarkRows32workers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Render(false, 32)
	}
}
func BenchmarkRows64workers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Render(false, 64)
	}
}

func BenchmarkSequential(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RenderSequential()
	}
}

func BenchmarkPixels(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Render(true, 12)
	}
}

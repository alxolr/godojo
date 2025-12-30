package solution

import "testing"

func BenchmarkMyFunction(b *testing.B) {
	bottom := "AAAA"
	allowed := []string{"AAB", "AAC", "BCD", "BBE", "DEF"}

	for i := 0; i < b.N; i++ {
		// Code to measure
		pyramidTransition(bottom, allowed)
	}
}

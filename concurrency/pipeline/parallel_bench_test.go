package pipeline

import "testing"

func BenchmarkMD5All(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testFunc(MD5All)
	}
}

func BenchmarkMD5AllWithCancel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testFunc(MD5All2)
	}
}

func BenchmarkMD5AllWithCancelAndBounded(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testFunc(MD5All3)
	}
}

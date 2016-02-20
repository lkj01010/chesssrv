package fw
import (
	"testing"
)

//go test -v -bench=".*"
//go test -bench=".*" -cpuprofile=cpu.prof -c

//func BenchmarkCryptoRand(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		Rand(1000)
//	}
//}

func BenchmarkFastRand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FastRand()
	}
}

func BenchmarkRandlibRand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Randn(1000000)
	}
}

func BenchmarkAdd(b *testing.B) {
	// 如果需要初始化，比较耗时的操作可以这样：
	// b.StopTimer()
	// .... 一堆操作
	// b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = 1 + 2
	}
}

func BenchmarkDiv(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = 12122 / 2221
	}
}
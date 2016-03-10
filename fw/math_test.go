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
func BenchmarkFastRandn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FastRandn(100000)
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

var vmap = make(map[int]int, 10000)
const maptimes = 100
func init(){
//	for i := 0; i < maptimes; i++ {
//		vmap[i] = i
//	}
}

// big map random access 30ns/op, write 230ns/op
func BenchmarkMapAccess(b *testing.B){
	temp := 0
	for i := 0; i < b.N; i ++ {
		vmap[i] = i
		temp = vmap[i]
		temp = vmap[i]
		temp = vmap[i]
		temp = vmap[i]
		temp = vmap[i]
		temp = vmap[i]
		temp = vmap[i]
	}
	b.Log(temp)
}
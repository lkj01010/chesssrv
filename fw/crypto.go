package fw
import (
//	"crypto/rand"
//	"math/big"
	"math/rand"
	"time"
)
// slowest 1898 ns/op
//func Rand(max int64) (ret int64) {
//	m := big.NewInt(max)
//	bigInt, _ := rand.Int(rand.Reader, m)
//	ret = bigInt.Int64()
//	return
//}

// 37.8 ns/op
func Randn(n int)int{
	return rand.Intn(n)
}

//----------------------
// from repo:gonet
var x0 uint32 = uint32(time.Now().UnixNano())
var a uint32 = 1664525
var c uint32 = 1013904223

// 3.0ns/op
func FastRand()int{
	x0 = a * x0 + c
	return int(x0)
}
//]]
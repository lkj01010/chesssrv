package fw
import (
	"crypto/rand"
	"math/big"
	"time"
)

func Rand(max int64) (ret int64) {
	m := big.NewInt(max)
	bigInt, _ := rand.Int(rand.Reader, m)
	ret = bigInt.Int64()
	return
}

//----------------------
// from repo:gonet
var x0 uint32 = uint32(time.Now().UnixNano())
var a uint32 = 1664525
var c uint32 = 1013904223

func FastRand()uint32{
	x0 = a * x0 + c
	return x0
}
//]]
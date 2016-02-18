package fw
import (
	"crypto/rand"
	"math/big"
)

func Rand(max int64) (ret int64) {
	m := big.NewInt(max)
	bigInt, _ := rand.Int(rand.Reader, m)
	ret = bigInt.Int64()
	return
}
package fw
import (
	"crypto/rand"
	"math/big"
)

func Rand(max int) (ret int) {
	m := big.NewInt(max)
	ret, _ = rand.Int(rand.Reader, m)
	return
}
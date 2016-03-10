package dao
import (
"testing"
"github.com/garyburd/redigo/redis"
"chess/cfg"
	"time"
	"fmt"
)

var redisConn redis.Conn
func init() {
	redisConn, _ = redis.Dial("tcp", cfg.RedisAddr(),
		redis.DialReadTimeout(1 * time.Second), redis.DialWriteTimeout(1 * time.Second))
	s, _ := redisConn.Do("SELECT", 10)
	fmt.Println("select 10, s=", s)

}

func BenchmarkDo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		redisConn.Do("GET", i, i)
	}
	b.Logf("BenchmarkDo, %+v times", b.N)
}

func BenchmarkGo(b *testing.B) {
//	b.Logf("redisConn=%+v", redisConn)
	acc := make(chan int, 0)
	do := 0
	for i := 0; i < b.N; i++ {
		go func() {
			redisConn.Do("GET", i, i)
			acc <- 1
			do ++
		}()
	}
	for i := 0; i <b.N; i++ {
		<- acc
	}
	b.Logf("BenchmarkGo, %+v times",do)
}
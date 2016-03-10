package dao
import (
"testing"
"github.com/garyburd/redigo/redis"
"chess/cfg"
	"time"
	"fmt"
	"sync"
)

var redisConn redis.Conn
func init() {
	redisConn, _ = redis.Dial("tcp", cfg.RedisAddr(),
		redis.DialReadTimeout(1 * time.Second), redis.DialWriteTimeout(1 * time.Second))
	s, _ := redisConn.Do("SELECT", 11)
	fmt.Println("select 11, s=", s)

}

//func BenchmarkDo(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		redisConn.Do("SET", i, i)
//	}
//	b.Logf("BenchmarkDo, %+v times", b.N)
//}

func BenchmarkGo(b *testing.B) {
//	b.Logf("redisConn=%+v", redisConn)
	acc := make(chan int, 0)
	do := 0
	mu := sync.Mutex{}
	for i := 0; i < b.N; i++ {
		go func() {
			mu.Lock()
			do = do + 1
			mu.Unlock()
			_, e := redisConn.Do("SET", do, "jjj")
			if e != nil {
				b.Log(e)
			}
			acc <- 1
		}()
	}
	for i := 0; i <b.N; i++ {
		<- acc
	}

//	time.Sleep(3 * time.Second)
	b.Logf("BenchmarkGo, %+v times",do)
}
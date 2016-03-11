package main

import (
	"github.com/garyburd/redigo/redis"
	"chess/cfg"
	"time"
	"fmt"
//	"sync"
	"github.com/lkj01010/log"
)

var redisConn redis.Conn
func init() {
}

//func BenchmarkDo(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		redisConn.Do("SET", i, i)
//	}
//	b.Logf("BenchmarkDo, %+v times", b.N)
//}

func main() {
	redisConn, _ = redis.Dial("tcp", cfg.RedisAddr(),
		redis.DialReadTimeout(1 * time.Second), redis.DialWriteTimeout(1 * time.Second))
	s, _ := redisConn.Do("SELECT", 11)
	fmt.Println("select 11, s=", s)

	defer redisConn.Close()


	N := 10000
//	acc := make(chan int, 0)
//	do := 0
//	mu := sync.Mutex{}
//
	t1 :=time.Now()

	for i := 0; i < N; i++ {
//		go func() {
//			mu.Lock()
//			do = do + 1
////			_, e := redisConn.Do("SET", do, "jjj")
			_, e := redisConn.Do("PING")
			if e != nil {
				fmt.Print(e, "\n")
			}
//			mu.Unlock()
//			acc <- 1
//		}()
	}
//	for i := 0; i < N; i++ {
//		<-acc
//	}

	t2 :=time.Now()
	d := t2.Sub(t1)

	//	time.Sleep(3 * time.Second)
//	log.Infof("BenchmarkGo, %+v times in %+v", do, d)
	log.Infof("BenchmarkGo, %+v times in %+v", N, d)
}


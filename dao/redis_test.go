package dao
import (
	"testing"
	"github.com/garyburd/redigo/redis"
//	"chess/cfg"
	"time"
	"fmt"
	"sync"
)

var redisConn redis.Conn
func init() {

//	c, err := redis.Dial("tcp", cfg.RedisAddr(),
//		redis.DialReadTimeout(1 * time.Second), redis.DialWriteTimeout(1 * time.Second))
	c, err := redis.Dial("tcp", fmt.Sprintf(":%d", 6379),
		redis.DialReadTimeout(1 * time.Second), redis.DialWriteTimeout(1 * time.Second))
	redisConn = c
	if err != nil {
		fmt.Print(err)
	}
	s, _ := redisConn.Do("SELECT", 11)
	fmt.Println("select 11, s=", s)

}

func zBenchmarkDoPing(b *testing.B) {
//	b.StopTimer()
	//	c, err := redis.DialDefaultServer()
	//lkj modify:
//	redisConn, err := redis.Dial("tcp", fmt.Sprintf(":%d", 6379), redis.DialReadTimeout(1 * time.Second), redis.DialWriteTimeout(1 * time.Second))
//	defer redisConn.Close()
//	b.StartTimer()
	for i := 0; i < b.N; i++ {
		if _, err := redisConn.Do("PING"); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkDo(b *testing.B) {
	b.StopTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		//		redisConn.Do("SET", i, i)
		//		redisConn.Do("PING")
		if _, err := redisConn.Do("PING"); err != nil {
			b.Fatal(err)
		}
	}
	//	b.Logf("BenchmarkDo, %+v times", b.N)
}

func BenchmarkGo(b *testing.B) {
	//	b.Logf("redisConn=%+v", redisConn)
	acc := make(chan int, 0)
	do := 0
	mu := sync.Mutex{}
	for i := 0; i < b.N; i++ {
		go func() {
			mu.Lock()
			do = do + 1
			_, e := redisConn.Do("SET", do, "jjj")
			if e != nil {
				b.Log(e)
			}
			mu.Unlock()
			acc <- 1
		}()
	}
	for i := 0; i < b.N; i++ {
		<-acc
	}

	//	time.Sleep(3 * time.Second)
	b.Logf("BenchmarkGo, %+v times", do)
}

func BenchmarkRedisPool(b *testing.B) {
	b.StopTimer()
	p := redis.Pool{
		Dial: func() (redis.Conn, error) {
//			c, err := redis.Dial("tcp", cfg.RedisAddr())
			c, err := redis.Dial("tcp", ":6379")
			if err != nil {
				return nil, err
			}
			// 选择db
			c.Do("SELECT", 11)
			return c, nil
		},

		MaxIdle: 30,
		MaxActive: 30}
	c := p.Get()
	if err := c.Err(); err != nil {
		b.Fatal(err)
	}
	c.Close()
	defer p.Close()
//	c = p.Get()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
				c = p.Get()
		if _, err := c.Do("PING"); err != nil {
			b.Fatal(err)
		}
				c.Close()
	}
	b.StopTimer()
	c.Close()

}
package cfg
import (
	"github.com/garyburd/redigo/redis"
	"fmt"
	"time"
)

func RedisAddr() (string) {
	addr := IPs[SrvType] + ":" + RedisPort
	return addr
}

func FlushCfgToDB() {
	c, err := redis.Dial("tcp", RedisAddr(), redis.DialReadTimeout(1 * time.Second), redis.DialWriteTimeout(1 * time.Second))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()

	c.Do("SELECT", RedisDBs["cfg"])
	c.Do("FLUSHALL")
	c.Do("SET", "srv:port:agent", AgentPort)
	c.Do("SET", "srv:port:lobby", LobbyPort)
	c.Do("SET", "srv:port:game", GamePort)
	c.Do("SET", "srv:port:data", DataPort)

	fmt.Println("FlushCfgToDB")
}
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
	c.Do("set", "AgentPort", AgentPort)
	c.Do("set", "LobbyPort", AgentPort)
	c.Do("set", "GamePort", AgentPort)
	c.Do("set", "DataPort", AgentPort)
}
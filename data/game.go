package data
import (
	"github.com/garyburd/redigo/redis"
	"chess/fw"
	"time"
"chess/cfg"
	"strconv"
)

type Game struct {
	c redis.Conn
}

func(this *Game)Init()(b bool, err error){
	//setup connection
	this.c, err = redis.Dial("tcp", cfg.RedisAddr(),
		redis.DialReadTimeout(1 * time.Second), redis.DialWriteTimeout(1 * time.Second))
	if err != nil {
		fw.Log.Error("data:game redis.Dial error")
		return
	}

	//select db
	_, err = this.c.Do("SELECT", cfg.RedisDBs["game"])
	if err != nil {
		fw.Log.Error("select err")
	}
	return
}

func(this *Game)Exit(){
	this.c.Close()
}

func(this *Game)GenLoginKey(id string)(key string) {
	fw.Log.Info("GenLoginKey")
	rand := fw.Rand(1000)
	key = strconv.FormatInt(rand, 10)
	this.c.Do("HSET", "loginkey", id, key)
	return
}
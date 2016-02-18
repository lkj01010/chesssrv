package data
import (
	"github.com/garyburd/redigo/redis"
	"chess/fw"
)

type Game struct {
	c redis.Conn
}

func(g *Game)prepare()(b bool, err error){
	return
}

func(g *Game)GenLoginKey(id string)(key string) {
	fw.Log.Info("GenLoginKey")
	rand := fw.Rand(100000)
	key = string(rand)
	g.c.Do("HSET", "loginkey", id, key)
	return
}

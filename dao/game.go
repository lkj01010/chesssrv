package dao
import (
	"github.com/garyburd/redigo/redis"
	"chess/fw"
	"time"
	"chess/cfg"
	"strconv"
)

const (
	LoginkeyKey = "loginkey"
)

type Game struct {
	c redis.Conn
}

var game *Game

func init() (b bool, err error) {
	//setup connection
	game = new(Game)
	game.c, err = redis.Dial("tcp", cfg.RedisAddr(),
		redis.DialReadTimeout(1 * time.Second), redis.DialWriteTimeout(1 * time.Second))
	if err != nil {
		fw.Log.Error("data:game redis.Dial error")
		return
	}

	//select db
	_, err = game.c.Do("SELECT", cfg.RedisDBs[cfg.Game])
	if err != nil {
		fw.Log.Error("select err")
	}
	return
}

func (g *Game)exit() {
	game.c.Close()
}

func (g *Game)genLoginKey(id string) (key string) {
	fw.Log.Info("game:genLoginKey")
	key = strconv.Itoa(fw.FastRand())
	g.c.Do("HSET", LoginkeyKey, id, key)
	return
}
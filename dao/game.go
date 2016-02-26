package dao
import (
	"github.com/garyburd/redigo/redis"
	"chess/fw"
	"time"
	"chess/cfg"
	"strconv"
	log "github.com/lkj01010/log"
)

const (
	LoginkeyKey = "loginkey"
)

type Game struct {
	c redis.Conn
}

var GameInst *Game

func NewGame() *Game{
	g := new(Game)
	c, err := redis.Dial("tcp", cfg.RedisAddr(),
		redis.DialReadTimeout(1 * time.Second), redis.DialWriteTimeout(1 * time.Second))
	if err != nil {
		log.Error(err.Error())
		return nil
	}
	g.c = c

	//select db
	_, err = c.Do("SELECT", cfg.RedisDBs[cfg.Game])
	if err != nil {
		log.Error(err.Error())
		return nil
	}

	return g
}

func init() {
	//setup connection
	GameInst = new(Game)
	var err error
	GameInst.c, err = redis.Dial("tcp", cfg.RedisAddr(),
		redis.DialReadTimeout(1 * time.Second), redis.DialWriteTimeout(1 * time.Second))
	if err != nil {
		log.Error("data:game redis.Dial error")
		return
	}

	//select db
	_, err = GameInst.c.Do("SELECT", cfg.RedisDBs[cfg.Game])
	if err != nil {
		log.Error(err.Error())
	}
	return
}

func (g *Game)exit() {
	GameInst.c.Close()
}

func (g *Game)genLoginKey(id string) (key string) {
	log.Info("game:genLoginKey")
	key = strconv.Itoa(fw.FastRand())
	g.c.Do("HSET", LoginkeyKey, id, key)
	return
}
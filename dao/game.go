package dao
import (
	"github.com/garyburd/redigo/redis"
	"chess/fw"
	"time"
	"chess/cfg"
	"strconv"
	log "github.com/lkj01010/log"
"chess/com"
)

const (
	LoginkeyKey = "loginkey"
)

type Game struct {
	model
}

func NewGame(p *Models) *Game {
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

	//register model
	g.parent = p

	return g
}

func (g *Game)exit() {
	g.c.Close()
}

func (g *Game)genLoginKey(id string) (key string) {
	log.Infof("game:genLoginKey, g=%+v", g)
	key = strconv.Itoa(fw.FastRand())
	log.Debugf("%+v, %+v", g, g.c)
	_, err := g.c.Do("HSET", LoginkeyKey, id, key)
	if err != nil {
		log.Error(err.Error())
	}
	return
}

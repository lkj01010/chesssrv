package dao
import "github.com/garyburd/redigo/redis"

// model
type model struct {
	c redis.Conn
	parent *Models
}

type Models struct {
	User *User
	Game *Game
}

func NewModels() *Models {
	m := new(Models)
	user := NewUser(m)
	game := NewGame(m)
	m.User = user
	m.Game = game
	return m
}

func (m *Models)Exit(){
	m.User.exit()
	m.Game.exit()
}
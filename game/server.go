package game
import (
"github.com/garyburd/redigo/redis"
)

type model struct {
c redis.Conn
parent *Models
}

type Models struct {
}

func NewModels() *Models {
m := new(Models)
return m
}

func (m *Models)Exit(){
}
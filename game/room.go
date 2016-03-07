package game
import (
	"sync"
	"chess/game/cow"
)

// player
type player struct {
	cmd chan string
}

// room
type Room struct {
	typ roomType

	mu sync.Mutex
	players map[string]*player
	game game
}

// useless
func NewRoom(typ roomType) *Room {

	return &Room{
		typ: typ,
		players: map[string]*player{},
//		game: cow.NewGame(c),
	}
}

func (r *Room)Serve() {
	c := make(chan string, 10)
	game := cow.NewGame(c)

	game.Serve()
}
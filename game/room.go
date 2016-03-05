package game
import (
	"sync"
	"chess/game/cow"
)

// player
type player struct {
	cmdChan chan string
}

// room
type roomType int
const (
	roomType_0 roomType = iota
)

type Room struct {
	typ roomType

	mu sync.Mutex
	players map[string]*player
	game game
}

func NewRoom(typ roomType) *Room {
	return &Room{
		typ: typ,
		players: map[string]*player{},
		game: cow.Game,
	}
}
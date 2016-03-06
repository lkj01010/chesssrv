package cow

type playerState int
const (
	psWait playerState = iota
	psPlay
)

type player struct {
	c     chan string
	id    string
	coin  int
	state playerState
	cards []card
}

func NewPlayer(c chan string, id string, coin int) *player {
	return &player{
		c: c,
		id: id,
		coin: coin,
		state: psWait,
		cards : make([]card, 5),
	}
}

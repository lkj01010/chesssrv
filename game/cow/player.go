package cow

type playerState int
const (
	psWait playerState = iota
	psPlay
)

type player struct {
	id    string
	coin  int
	state playerState
	cards []card
}

func NewPlayer(id string, coin int) *player {
	return &player{
		id: id,
		coin: coin,
		state: psWait,
		cards : make([]card, 5),
	}
}

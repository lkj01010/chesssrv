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
	sendFunc func(msg string)
}

func NewPlayer(id string, sendFunc func(msg string)) *player {
	return &player{
		id: id,
		state: psWait,
		cards : make([]card, 5),
		sendFunc: sendFunc,
	}
}

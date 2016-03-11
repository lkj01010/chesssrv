package cow

type playerState int
const (
	psWait playerState = iota
	psPlay
)

type player struct {
	id       string
	state    playerState
	cards    []card

	rcvr     chan string
	sendFunc func(msg string)

	//player info
	coin     int
}

func NewPlayer(id string, rcvr chan string, sendFunc func(msg string)) *player {
	return &player{
		id: id,
		state: psWait,
		cards : make([]card, 5),
		rcvr: rcvr,
		sendFunc: sendFunc,
	}
}

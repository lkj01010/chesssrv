package cow

type playerState int
const (
	psWait playerState = iota
	psPlay
)

type player struct {
	id string
	coin  int
	state playerState
}
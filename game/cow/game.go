package cow

type DaoCtrl interface {
	AddCoin(id int, coin int)

}

type gameState int
const (
	gsDeal gameState = iota    // 发牌
	gsPlay
	gsSettle    // 结算
)

type Game struct {
	DaoCtrl
	c       chan string
	players []player
}

func NewGame(c chan string) *Game {
	return &Game{
		c: c,
		players: []player{},
	}
}
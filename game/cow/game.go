package cow
import (
	"time"
	"github.com/lkj01010/log"
)

type DaoCtrl interface {
	AddCoin(id int, coin int)
}

type gameState int
const (
	gsWait gameState = iota    // 发牌
	gsDeal
	gsPlay
	gsSettle    // 结算
)

type Game struct {
	DaoCtrl
	c       chan string
	state   gameState
	players []*player
	timer   *time.Timer
}

func NewGame(c chan string) *Game {
	return &Game{
		c: c,
		players: make([]*player, 0, maxPlayer),
		timer: time.NewTimer(0),
	}
}

func (g *Game)Serve() {
	for {
		select {
		case <-g.timer.C:
			g.onTimer()
		case msg := <-g.c:
			log.Debugf("game chan recv=%+v", msg)
		}
	}
}

func (g *Game)PlayerEnter(c chan string, id string, coin int) {
	g.players = append(g.players, NewPlayer(c, id, coin))
}

func (g *Game)NewRound() {
	g.Deal()
	g.timer.Reset(timeout_settle)
}

func (g* Game)Deal() {
	for _, player := range (g.players) {
		DealCards(player.cards)
	}
}

func (g *Game)Settle() {

}

func (g *Game)onTimer() {

}
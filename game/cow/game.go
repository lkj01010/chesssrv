package cow
import (
	"time"
	"github.com/lkj01010/log"
)

type DaoCtrl interface {
	AddCoin(id string, coin int)
	GetMultiCoin(id []string) map[string]int
}

type gameState int
const (
	gsWait gameState = iota    // 发牌
	gsDeal
	gsSettle    // 结算
)

type Game struct {
	DaoCtrl
	c        chan string
//	roomType RoomType
	enterCoin int

	state    gameState
	players  []*player
	timer    *time.Timer
}

func NewGame(c chan string, enterCoin int) *Game {
	return &Game{
		c: c,
		enterCoin: enterCoin,
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

func (g *Game)PlayerEnter(id string, coin int) {
	g.players = append(g.players, NewPlayer(id, coin))
}

func (g *Game)onTimer() {
	switch g.state {
	case gsWait:
		g.Deal()
	case gsDeal:
		g.settle()
	case gsSettle:
		g.NewRound()
	}
}

func (g *Game)NewRound() {
	g.state = gsWait
	// 检查每个玩家钱是否足够
	g.checkCoinEnough()
	// 等待玩家准备就绪
	g.timer.Reset(timeout_deal)
}

func (g* Game)Deal() {
	g.state = gsDeal
	for _, player := range (g.players) {
		DealCards(player.cards)
	}
	g.timer.Reset(timeout_settle)
}

func (g *Game)checkCoinEnough() {
	ids := []string{}
	for _, player := range (g.players) {
		if (player.state == psPlay) {
			ids = append(ids, player.id)
		}
	}
	playerCoinMap := g.GetMultiCoin(ids)

//	for id, coin := range (playerCoinMap) {
	for _, coin := range (playerCoinMap) {
		if coin < g.enterCoin {
			// 没有足够的钱玩
//			g.c <- msgcInst.hasNoEnoughMoney(id)
			// todo: connId
			g.c <- msgcInst.hasNoEnoughMoney(0)
		}
	}
}

func (g *Game)settle() {
	g.timer.Reset(timeout_newrount)
}
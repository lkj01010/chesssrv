package cow
import (
	"time"
	"sync"
	"net/rpc"
	"chess/dao"
)

type gameState int
const (
	gsWait gameState = iota    // 发牌
	gsDeal
	gsSettle    // 结算
	gsExit    // 全部人退出,销毁
)

type PlayerMsg struct {
	Id  string
	msg string
}

type Game struct {
	dao       *rpc.Client
	id        int
	//	roomType RoomType
	enterCoin int

	state     gameState

	players   []*player
	playersMu sync.Mutex

	timer     *time.Timer
}

func NewGame(dao *rpc.Client, id, enterCoin int) *Game {
	return &Game{
		dao: dao,
		id: id,
		enterCoin: enterCoin,
		players: make([]*player, 0, maxPlayer),
		timer: time.NewTimer(0),
	}
}

func (g *Game)Go() {
	for {
		select {
		case <-g.timer.C:
			g.onTimer()
		}

		for i, player := range g.players {
			select {
			case msg := <-player.rcvr:
				g.handlePlayerMsg(i, msg)
			}
		}

		if g.state == gsExit {
			return
		}
	}
}

func (g *Game)PlayerEnter(id string, rcvr chan string, sendFunc func(msg string)) {
	player := NewPlayer(id, rcvr, sendFunc)

	g.playersMu.Lock()
	g.players = append(g.players, player)
	g.playersMu.Unlock()

	// get info from db
	var reply dao.Reply
	g.dao.Call("User.GetCoin", &dao.Args{Id: id}, &reply)

	player.coin = reply.Int
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
	for _, player := range (g.players) {
		if player.state == psPlay && player.coin < g.enterCoin {
			player.sendFunc(msgCreatorInst.hasNoEnoughMoney())
		}
	}
}

func (g *Game)settle() {
	g.timer.Reset(timeout_newrount)
}

func (g *Game)handlePlayerMsg(idx int, msg string) {

}
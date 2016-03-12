package cow
import (
	"time"
	"sync"
	"net/rpc"
	"chess/dao"
	"fmt"
	"chess/com"
	"encoding/json"
)

type gameState int
const (
	gsWaitPlayer gameState = iota    // 发牌
	gsWaitReady
	gsDealAndPlay
	gsSettle    // 结算
	gsExit    // 全部人退出,销毁
)

const cmd_PlayerEnter = 500000

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
	playersMu sync.RWMutex

	timer     *time.Timer
}

func NewGame(dao *rpc.Client, id, enterCoin int) *Game {
	return &Game{
		dao: dao,
		id: id,
		enterCoin: enterCoin,
		state: gsWaitPlayer,
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

		g.playersMu.RLock()
		for i, player := range g.players {
			select {
			case msg := <-player.rcvr:
				g.handlePlayer(i, msg)
			}
		}
		g.playersMu.RUnlock()

		if g.state == gsExit {
			return
		}
	}
}

func (g *Game)PlayerEnter(id string, rcvr chan string, sendFunc func(msg string)) {
	player := NewPlayer(id, rcvr, sendFunc)

	// get info from db
	var reply dao.User_InfoReply
	g.dao.Call("User.GetInfo", &dao.Args{Id: id}, &reply)

	player.info = reply.Info

	// add to game
	g.playersMu.Lock()
	g.players = append(g.players, player)
	g.playersMu.Unlock()

	// notify game "i come"
	rcvr <-com.MakeMsgString(cmd_PlayerEnter, 0, nil)
}

func (g *Game)onTimer() {
	switch g.state {
	case gsWaitPlayer:
		g.Deal()
	case gsWaitReady:
	case gsDealAndPlay:
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
	for _, player := range (g.players) {
		// test
		player.sendFunc(msgCreatorInst.Settle(&SettleNtf{TempString:fmt.Sprintf("id=%+v, string=%+v", player.id, "haha")}))
	}
	g.timer.Reset(timeout_newrount)
}

func (g *Game)handlePlayer(idx int, msgstr string) {
	var msg com.Msg
	if err := json.Unmarshal([]byte(msgstr), &msg); err != nil {
		return
	}

	switch msg.Cmd {
	case cmd_PlayerEnter:
		handlePlayerEnter(idx)
	}

}

func (g *Game)handlePlayerEnter(idx int){
	// notify
	msg := msgCreatorInst.PlayerEnter(i)
}
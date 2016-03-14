package cow
import (
	"time"
	"net/rpc"
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

type PlayerEnterOrLeave struct {
	isEnter  bool
	id       string
	info     com.UserInfo
	rcvr     chan string
	sendFunc func(msg string)
}

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
	playerCh  chan PlayerEnterOrLeave
	//			  playersMu sync.RWMutex

	timer     *time.Timer
}

func NewGame(dao *rpc.Client, id, enterCoin int) *Game {
	return &Game{
		dao: dao,
		id: id,
		enterCoin: enterCoin,
		state: gsWaitPlayer,
		players: make([]*player, 0, maxPlayer),
		playerCh: make(chan PlayerEnterOrLeave, 10),
		timer: time.NewTimer(0),
	}
}

func (g *Game)Go() {
	for {
		select {
		case el := <-g.playerCh:
			g.handlePlayerEnterOrLeave(el)

		case <-g.timer.C:
			g.onTimer()
		}

		for i, player := range g.players {
			select {
			case msg := <-player.rcvr:
				g.handlePlayerReceive(i, msg)
			}
		}

		if g.state == gsExit {
			return
		}
	}
}

func (g *Game)handlePlayerEnterOrLeave(el PlayerEnterOrLeave) {
	if el.isEnter {
		player := NewPlayer(el.id, el.info, el.rcvr, el.sendFunc)
		g.players = append(g.players, player)

		// notify game "i come"
		el.rcvr <- com.MakeMsgString(com.Cmd_Game_EnterReq, 0, nil)

	} else {
		// todo:
	}
}

func (g *Game)PlayerEnter(id string, info com.UserInfo, rcvr chan string, sendFunc func(msg string)) {
	g.playerCh <- PlayerEnterOrLeave{
		id: id,
		info: info,
		rcvr: rcvr,
		sendFunc: sendFunc,
	}
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
	g.state = gsWaitReady
	// 检查每个玩家钱是否足够
	g.checkCoinEnough()
	// 等待玩家准备就绪
	g.timer.Reset(timeout_deal)
}

func (g* Game)Deal() {
	g.state = gsDealAndPlay
	for _, player := range (g.players) {
		DealCards(player.cards)
	}
	g.timer.Reset(timeout_settle)
}

func (g *Game)checkCoinEnough() {
	for _, player := range (g.players) {
		if player.state == psPlay && player.info.Coin < g.enterCoin {
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

func (g *Game)handlePlayerReceive(idx int, msgstr string) {
	var msg com.Msg
	if err := json.Unmarshal([]byte(msgstr), &msg); err != nil {
		return
	}

	switch com.Command(msg.Cmd) {
	case com.Cmd_Game_EnterReq:
		g.handlePlayerEnter(idx)
	}

}

func (g *Game)handlePlayerEnter(idx int) {
	// notify
	msg := msgCreatorInst.PlayerEnter(idx, &g.players[idx].info)
	for _, player := range g.players {
		player.sendFunc(msg)
	}

	if g.state == gsWaitPlayer {
		// check game state
		if len(g.players) >= 2 {
			g.state = gsWaitReady

			for _, player := range g.players {
				msg := msgCreatorInst.WaitReady()
				player.sendFunc(msg)
			}
		}
	}


}
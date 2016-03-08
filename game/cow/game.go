package cow
import (
	"time"
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
	roomType RoomType
	state    gameState
	players  []*player
	timer    *time.Timer
}

func NewGame(c chan string, rt RoomType) *Game {
	return &Game{
		c: c,
		roomType: rt,
	}
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

	for id, coin := range (playerCoinMap) {
		if coin < RoomEnterCoin[g.roomType] {
			// 没有足够的钱玩
			g.c <- msgcInst.hasNoEnoughMoney(id)
		}
	}
}

func (g *Game)settle() {
	g.timer.Reset(timeout_newrount)
}
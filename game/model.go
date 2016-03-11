package game
import (
	"net/rpc"
	"github.com/lkj01010/log"
	"chess/cfg"
	"time"
	"chess/fw"
	"chess/game/cow"
	"encoding/json"
	"chess/com"
	"sync"
)

var modelInst *Model

type Model struct {
	dao          *rpc.Client

	agent        *fw.Agent             // 持有的外部agent (用来发消息)

	playerAgents map[int]*playerAgent  // 消息分发到用户的agent
	mu           sync.Mutex

	gameIdAcc    int                   // 房间id counter

	games        map[int]*cow.Game     // 房间
	freeGames    [RoomType][]*cow.Game // 每个type一个数组
	gamesMu      sync.RWMutex

	playerGames  map[string]*cow.Game  // 玩家（id） 对应其进入的game
}

func NewModel() *Model {
	if modelInst != nil {
		panic("game:Model:NewModel can invoke only once")
	}
	modelInst := new(Model)

	// connect dao server
	dao:    daocli, err := rpc.Dial("tcp", cfg.DaoAddr())
	if err != nil {
		log.Warningf("dao server connect fail(err=%+v), try again...", err.Error())
		time.Sleep(1 * time.Second)
		goto dao
	}

	modelInst.dao = daocli

	// data init
	modelInst.gameIdAcc = 0
	modelInst.games = map[int]*cow.Game{}
	modelInst.playerGames = map[string]*cow.Game{}
	modelInst.freeGames = [][]*cow.Game{}
	// 1000 games each type
	for i, _ := range (modelInst.freeGames) {
		modelInst.freeGames[i] = make([]*cow.Game, 0, 1000)
	}

	return modelInst
}

func (m *Model)Hook(a *fw.Agent) {
	m.agent = a
}

func (m *Model)Enter() {

}

func (m *Model)Exit() {
	m.dao.Close()
}

func (m *Model)playerAgentSend(connId int, msg string) {
	idmsg := com.MakeConnIdRawMsgString(connId, msg)
	m.agent.Send(idmsg)
}

func (m *Model)Handle(req string) (resp string, err error) {
	var outmsg com.ConnIdRawMsg
	if err = json.Unmarshal([]byte(req), &outmsg); err != nil {
		return
	}
	connId := outmsg.ConnId

	// find agent
	var pa *playerAgent
	pa, err = m.playerAgents[connId]
	if err == nil {

	} else {
		pa = NewPlayerAgent(connId, m.playerAgentSend)
		pa.Go()
		m.playerAgents[connId] = pa
	}

	pa.Receive(outmsg.Content)
	return
}

func (m *Model)GetFreeGameByType(typ RoomType) (game *cow.Game) {
	games := modelInst.freeGames[typ]

	m.gamesMu.Lock()
	defer m.gamesMu.Unlock()

	var game *cow.Game
	if len(games) == 0 {
		// create game
		m.gameIdAcc++
		game = cow.NewGame(m.dao, m.gameIdAcc, RoomEnterCoin[typ])
		game.Go()

		games = append(games, game)
		m.games[m.gameIdAcc] = game

	} else {
		// put to first game
		game = games[0]
	}
	return
}

func (m *Model)RemovePlayerAgent(connId string) {

}
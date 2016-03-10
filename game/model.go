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

	// 持有的外部agent (用来发消息)
	agent        *fw.Agent

	// 消息分发到用户的agent
	playerAgents map[int]*playerAgent
	mu           sync.Mutex

	// 房间id counter
	roomIdAcc    int
	// 房间
	games        map[int]*cow.Game
	// 玩家（id） 对应其进入的game
	playerGames  map[string]*cow.Game
	// 每个type一个数组
	freeGames    [RoomType][]*cow.Game
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
	modelInst.roomIdAcc = 0
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

func (m *Model)playerAgentSend(connId int, msg string){
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





	pa.Recive(outmsg.Content)


	if msg.Cmd == Cmd_Game_EnterReq {
		_, e := m.playerGames[connId]
		if e == nil {
			log.Warning("game:model:handle:player enter req, err=already in game")
			return com.ErrAlreadyInGame
		}


	} else {

	}




}


///////////////////////////////////////////////////////
// enter
func (m *Model)enterGame(id string, roomType RoomType) {

}
////////////////////////////////////////////
func (m *Model)handleEnter(connId, content string) (err error) {
	var req EnterGame
	if err = json.Unmarshal([]byte(content), &req); err != nil {
		log.Error("content=", content, ", err: ", err.Error())
		return
	}

	m.enterGame(connId, )
}

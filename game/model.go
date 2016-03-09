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
)

type Model struct {
	dao         *rpc.Client

	// 持有的外部agent (用来发消息)
	agent       *fw.Agent
	// 房间id counter
	roomId      int
	// 房间
	games       map[int]*cow.Game
	// 玩家（id） 对应其进入的game
	playerGames map[string]*cow.Game
	// 每个type一个数组
	freeGames   [][]*cow.Game
}

func NewModel() *Model {
	s := new(Model)

	// connect dao server
	dao:    daocli, err := rpc.Dial("tcp", cfg.DaoAddr())
	if err != nil {
		log.Warningf("dao server connect fail(err=%+v), try again...", err.Error())
		time.Sleep(1 * time.Second)
		goto dao
	}

	s.dao = daocli

	// data init
	s.roomId = 0
	s.games = map[int]*cow.Game{}
	s.playerGames = map[string]*cow.Game{}
	s.freeGames = [][]*cow.Game{}
	// 1000 games each type
	for i, _ := range (s.freeGames) {
		s.freeGames[i] = make([]*cow.Game, 0, 1000)
	}

	return s
}

func (m *Model)Hook(a *fw.Agent) {
	m.agent = a
}

func (m *Model)Enter() {

}

func (m *Model)Exit() {
	m.dao.Close()
}

func (m *Model)Handle(req string) (resp string, err error) {
	var outmsg com.ConnIdRawMsg
	if err = json.Unmarshal([]byte(req), &outmsg); err != nil {
		return
	}

	var msg com.Msg
	if err = json.Unmarshal([]byte(req), &msg); err != nil {
		return
	}


	switch msg.Cmd {
	case Cmd_Game_EnterReq:
		err = m.handleEnter(msg.Content)
		return
	}
}

///////////////////////////////////////////////////////
// enter
func (m *Model)enterGame() {

}
////////////////////////////////////////////
func (m *Model)handleEnter(content string) (err error) {
	var req EnterGame
	if err = json.Unmarshal([]byte(content), &req); err != nil {
		log.Error("content=", content, ", err: ", err.Error())
		return
	}
	rt := req.RoomType

}

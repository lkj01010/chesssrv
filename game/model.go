package game
import (
	"net/rpc"
	"github.com/lkj01010/log"
	"chess/cfg"
	"time"
	"chess/fw"
	"chess/game/cow"
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

	return
}
////////////////////////////////////////////
type Game_EnterArgs struct {

}

type Game_EnterReply struct {

}

func (s *Model)EnterGame(args *Game_EnterArgs, reply *Game_EnterReply) error {

	return nil
}

////////////////////////////////////////////
type Game_LeaveArgs struct {

}
type Game_LeaveReply struct {

}

func (s *Model)LeaveGame(args*Game_LeaveArgs, reply *Game_LeaveReply) error {
	return nil
}

////////////////////////////////////////////


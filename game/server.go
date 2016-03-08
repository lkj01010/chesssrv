package game
import (
	"net/rpc"
	"github.com/lkj01010/log"
	"chess/cfg"
	"time"
	"chess/game/cow"
)

// 封装一层是因为直接调用Server做rpc导出类,它有Close这个函数,不符合rpc类规范,报警告
type Model struct {
	Server *Server
}

func NewModel() *Model {
	return &Model{Server: NewServer()}
}

func (m *Model)Exit() {
	m.Server.close()
}

///////////////////////////////////////////////////////
type Server struct {
	dao         *rpc.Client

	// 房间id counter
	roomId      int
	// 房间
	games       map[int]*cow.Game
	// 玩家（id） 对应其进入的game
	playerGames map[string]*cow.Game
	// 每个type一个数组
	freeGames   [][]*cow.Game
}

func NewServer() *Server {
	s := new(Server)

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
		s.games[i] = make([]*cow.Game)
	}

	return s
}

func (s *Server)close() {
	if err := s.dao.Close(); err != nil {
		log.Error(err.Error())
	}
}

////////////////////////////////////////////
type Game_EnterArgs struct {

}

type Game_EnterReply struct {

}

func (s *Server)EnterGame(args *Game_EnterArgs, reply *Game_EnterReply) error {

	return nil
}

////////////////////////////////////////////
type Game_LeaveArgs struct {

}
type Game_LeaveReply struct {

}

func (s *Server)LeaveGame(args*Game_LeaveArgs, reply *Game_LeaveReply) error {
	return nil
}

////////////////////////////////////////////


package game
import (
	"net/rpc"
	"github.com/lkj01010/log"
	"chess/cfg"
	"time"
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
	dao   *rpc.Client

	rooms [][]*Room
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

	// 1000 rooms each type
	s.rooms = make([][]*Room, roomTypeCount)
	for i, _ := range (s.rooms) {
		s.rooms[i] = make([]*Room, 1000)
	}
	//	s.rooms = make(make([]*Room, 0, 1000), roomTypeCount)

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


package game
import (
	"net/rpc"
)

type Server struct {
	Dao    *rpc.Client

	rooms	[][]*Room
}

func (s *Server)init() {

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
func (s *Server)LeaveGame(argbbbGame_LeaveArgs, reply *Game_LeaveReply) error {
	return nil
}

////////////////////////////////////////////


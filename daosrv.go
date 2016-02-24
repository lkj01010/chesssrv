package main
import (
	"net/rpc"
	"net"
	"chess/cfg"
	"chess/dao"
	"github.com/lkj01010/log"
)

func main() {
	defer func() {
		dao.Exit()
	}()

	//register rpc
	rpc.Register(dao.UserInst)
	rpc.Register(dao.GameInst)

	//network
	serverAddr := "127.0.0.1:" + cfg.DaoPort
	l, e := net.Listen("tcp", serverAddr) // any available address
	if e != nil {
		log.Fatalf("net.Listen tcp : %v", e)
	}
	log.Infof("dao RPC server listening on", serverAddr)
	rpc.Accept(l)
}
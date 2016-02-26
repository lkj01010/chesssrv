package main
import (
	"net/rpc"
	"net"
	"chess/cfg"
	"chess/dao"
	"github.com/lkj01010/log"
)

func main() {
	//register rpc
	models := dao.NewModels()
	rpc.Register(models.User)
	rpc.Register(models.Game)
	defer func() {
		models.Exit()
	}()

	//network
	serverAddr := "127.0.0.1:" + cfg.DaoPort
	l, e := net.Listen("tcp", serverAddr) // any available address
	if e != nil {
		log.Fatalf("net.Listen tcp : %v", e)
	}
	log.Info("dao RPC server listening on ", serverAddr)
	rpc.Accept(l)
}
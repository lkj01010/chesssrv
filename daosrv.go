package main
import (
	"net/rpc"
	"chess/fw"
	"net"
	"chess/cfg"
	"chess/dao"
)

func main() {
	defer func(){
		dao.Exit()
	}()

	//register rpc
	rpc.Register(dao.UserInst)
	//network
	serverAddr := "127.0.0.1:" + cfg.DaoPort
	l, e := net.Listen("tcp", serverAddr) // any available address
	if e != nil {
		fw.Log.Fatalf("net.Listen tcp :0: %v", e)
	}
	fw.Log.Println("dao RPC server listening on", serverAddr)
	rpc.Accept(l)
}
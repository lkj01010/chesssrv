package dao

import (
	"testing"
	"net/rpc"
	"chess/cfg"
	log "github.com/lkj01010/log"
	"net"
	"sync"
)

var once sync.Once

func startServer() {
	//register rpc
	models := NewModels()
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
	go rpc.Accept(l)
}

func TestBareClient(t *testing.T) {
	once.Do(func() {
		go startServer()
	})

	client, err := rpc.Dial("tcp", "127.0.0.1:" + cfg.DaoPort)
	if err != nil {
		log.Error("dialing:", err)
	}

	log.Info("TestClient")

	// register
	{
		args := &User_RegisterArgs{"xxxaccount", "xxxpsw"}
		var reply RpcReply
		err = client.Call("User.HandleRegister", args, &reply)
		if err != nil {
			log.Error("HandleRegister error:", err)
		}
		log.Infof("HandleRegister: %v %v -> %+v", args.Account, args.Psw, reply)
	}

	// auth
	{
		args := &User_AuthArgs{"xxxaccount", "xxxpsw"}
		var reply User_AuthReply
		if err = client.Call("User.HandleAuth", args, &reply); err != nil {
			log.Error("HandleAuth", err)
		}
		log.Infof("HandleAuth: %v %v -> %+v", args.Account, args.Psw, reply)
	}
}

func TestDaoClient(t *testing.T) {
	//	once.Do(func() {go startServer()})
	once.Do(startServer)

	//	time.Sleep(3 * time.Second)

	cli, err := NewClient()
	log.Debug("tdc 1")
	if err != nil {
		log.Error("new client error: ", err.Error())
	}
	{
		var reply RpcReply
		args := User_AuthArgs{"zhu001", "21882"}
		if err := cli.UserRegister(&args, &reply); err != nil {
			log.Error(err.Error())
		}
	}
	{
		var reply User_AuthReply
		args := User_AuthArgs{"zhu001", "21882"}
		if err := cli.UserAuth(&args, &reply); err != nil {
			log.Error(err.Error())
		}
	}
}
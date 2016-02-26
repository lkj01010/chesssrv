package dao

import (
	"testing"
	"chess/fw"
	"net/rpc"
	"chess/cfg"
	log "github.com/lkj01010/log"
)

func testLogin(t *testing.T) {
	//	defer func(){
	//		Exit()
	//	}()

	//
	var registerReply fw.RpcReply
	UserInst.HandleRegister(&UserRegisterArgs{Account:"lkj2", Psw:"lkjpassword"}, &registerReply)

	var reply UserAuthReply
	UserInst.HandleAuth(&UserAuthArgs{Account:"lkj2", Psw:"lkjpassword"}, &reply)
}

func TestClient(t *testing.T) {
	client, err := rpc.Dial("tcp", "127.0.0.1:" + cfg.DaoPort)
	if err != nil {
		log.Error("dialing:", err)
	}

	//	log.SetFlags(log.LstdFlags|log.Lshortfile)
	//	log.SetPrefix("[DEBU]")
	//	log.Print("xxxx")

	log.Info("xjxjxjxj")

	// register
	{
		args := &UserRegisterArgs{"xxxaccount", "xxxpsw"}
		var reply fw.RpcReply
		err = client.Call("User.HandleRegister", args, &reply)
		if err != nil {
			log.Error("HandleRegister error:", err)
		}
		log.Infof("HandleRegister: %v %v -> %+v", args.Account, args.Psw, reply)
	}

	// auth
	{
		args := &UserAuthArgs{"xxxaccount", "xxxpsw"}
		var reply UserAuthReply
		if err = client.Call("User.HandleAuth", args, &reply); err != nil {
			log.Error("HandleAuth", err)
		}
		log.Infof("HandleAuth: %v %v -> %+v", args.Account, args.Psw, reply)
	}

}
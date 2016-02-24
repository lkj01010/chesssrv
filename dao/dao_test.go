package dao

import (
	"testing"
	"chess/fw"
	"net/rpc"
	"chess/cfg"
)

func testLogin(t *testing.T){
//	defer func(){
//		Exit()
//	}()

	//
	var registerReply fw.RpcReply
	UserInst.HandleRegister(&RegisterArgs{Account:"lkj2", Psw:"lkjpassword"}, &registerReply)

	var reply AuthReply
	UserInst.HandleAuth(&AuthArgs{Account:"lkj2", Psw:"lkjpassword"}, &reply)
}

func TestClient(t *testing.T){
	fw.Log.Debug("begin")
	client, err := rpc.Dial("tcp", "127.0.0.1:" + cfg.DaoPort)
	if err != nil {
		fw.Log.Fatal("dialing:", err)
	}
	fw.Log.Debug("conn ok")
	args := &RegisterArgs{"xxxaccount","xxxpsw"}
	var reply fw.RpcReply
	err = client.Call("User.HandleRegister", args, &reply)
	if err != nil {
		fw.Log.Fatal("arith error:", err)
	}
	fw.Log.Printf("HandleRegister: %v %v -> %+v", args.Account, args.Psw, reply)
}
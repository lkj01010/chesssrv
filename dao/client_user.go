package dao
import (
//	"github.com/lkj01010/log"
)

//func (cli *Client)UserRegister(args *User_RegisterArgs, reply *RpcReply) (err error) {
//	err = cli.Call("User.HandleRegister", args, &reply)
//	if err != nil {
//		log.Debug("call failed: ", err.Error())
//		return
//	}
//	log.Infof("UserRegister: %v %v -> %+v", args.Account, args.Psw, reply)
//	return
//}
//
//func (cli *Client)UserAuth(args *User_AuthArgs, reply *User_AuthReply) (err error) {
//	if err = cli.Call("User.HandleAuth", args, &reply); err != nil {
//		return
//	}
//	log.Infof("UserAuth: %+v -> %+v", args, reply)
//	return
//}
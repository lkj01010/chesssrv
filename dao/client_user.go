package dao
import (
"chess/fw"
"github.com/lkj01010/log"
)

func (cli *Client)UserRegister(accout, psw string, reply *fw.RpcReply) (err error){
	args := &RegisterArgs{accout, psw}
	var reply fw.RpcReply
	err = cli.cli.Call("User.HandleRegister", args, &reply)
	if err != nil {
		return
	}
	log.Infof("UserRegister: %v %v -> %+v", args.Account, args.Psw, reply)
	return
}

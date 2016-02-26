package dao
import (
	"chess/fw"
	"github.com/lkj01010/log"
)

func (cli *Client)UserRegister(account, psw string, reply *fw.RpcReply) (err error) {
	args := &UserRegisterArgs{account, psw}
	err = cli.cli.Call("User.HandleRegister", args, &reply)
	if err != nil {
		return
	}
	log.Infof("UserRegister: %v %v -> %+v", args.Account, args.Psw, reply)
	return
}

func (cli *Client)UserAuth(account, psw string, reply *UserAuthReply) (err error) {
	args := &UserAuthArgs{account, psw}
	if err = cli.cli.Call("User.HandleAuth", args, &reply); err != nil {
		return
	}
	log.Infof("UserAuth: %+v -> %+v", args, reply)
	return
}
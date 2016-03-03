package agent
import (
	"encoding/json"
	"chess/com"
	"chess/dao"
	log "github.com/lkj01010/log"
	"net/rpc"
)

//数据处理模块
type model struct {
	dao     *rpc.Client

	isLogin bool
	id      string

}

func NewModel(dao *rpc.Client, isTimeout bool) *model {
	m := &model{
		dao: dao,
	}
	return m
}

func (m *model)Enter() {
	m.isLogin = false
	m.id = "undefined"
}

func (m *model)Exit() {
}

func (m *model)Handle(req string) (resp string, err error) {
	var msg com.Msg
	if err = json.Unmarshal([]byte(req), &msg); err != nil {
		log.Error("Unmarshal err: ", err)
		return
	}

	log.Debugf("msg: %#v", req)

	switch msg.Cmd {
	case cmdHeartbeat:
		m.handleHeartbeat()
	case cmdRegisterReq:
		resp, err = m.handleRegister(msg.Content)
	case cmdAuthReq:
		resp, err = m.handleAuth(msg.Content)
	case cmdLoginReq:
		resp, err = m.handleLogin(msg.Content)
	}

	if err != nil {
		log.Error("handle err: ", err.Error())
	}
	return
}
func (m *model)handleHeartbeat() {
	// do nothing
}

func (m *model)handleRegister(content string) (resp string, err error) {
	//	daocli.
	var req RegisterReq
	if err = json.Unmarshal([]byte(content), &req); err != nil {
		log.Error("content=", content, ", err: ", err.Error())
		return
	}
	args := &dao.User_RegisterArgs{req.Account, req.Psw}
	var reply dao.RpcReply
	log.Debugf("req : %#v", req)
	//	if err = h.dc.UserRegister(&args, &reply); err != nil {
	if err = m.dao.Call("User.Register", args, &reply); err != nil {
		return
	}
	log.Infof("User.Register %+v -> %+v", args, reply)
	resp = com.MakeMsgString(cmdRegisterResp, reply.Code, nil)
	return
}

func (m *model)handleAuth(content string) (resp string, err error) {
	var req AuthReq
	if err = json.Unmarshal([]byte(content), &req); err != nil {
		return
	}
	args := &dao.User_AuthArgs{req.Account, req.Psw}
	var reply dao.User_AuthReply
	if err = m.dao.Call("User.Auth", args, &reply); err != nil {
		return
	}
	resp = com.MakeMsgString(cmdAuthResp, reply.Code, nil)
	return
}

func (m *model)handleLogin(content string) (resp string, err error) {
	var req LoginReq
	if err = json.Unmarshal([]byte(content), &req); err != nil {
		return
	}
	args := &dao.User_AuthArgs{req.Account, req.Psw}
	var reply dao.User_AuthReply
	if err = m.dao.Call("User.Auth", args, &reply); err != nil {
		return
	}
	resp = com.MakeMsgString(cmdLoginResp, reply.Code, nil)
	return
}
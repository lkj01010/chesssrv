package agent
import (
	"encoding/json"
	"chess/com"
	"chess/dao"
	log "github.com/lkj01010/log"
	"time"
	"net/rpc"
)
const (
	timeoutDuration = 5 * time.Second
)

type handler struct {
	dao            *rpc.Client
	isTimeout      bool
	heartbeatTimer time.Timer
}

func NewHandle(dao *rpc.Client, isTimeout bool) *handler {
	t := time.NewTimer(timeoutDuration)

}

func (h *handler)Handle(req string) (resp string, err error) {
	var msg com.Msg
	if err = json.Unmarshal([]byte(req), &msg); err != nil {
		log.Error("Unmarshal err: ", err)
		return
	}

	log.Debugf("msg: %#v", req)

	switch msg.Cmd {
	case cmdHeartbeat:
		resp, err = h.handleHeartbeat()
	case cmdRegisterReq:
		resp, err = h.handleRegister(msg.Content)
	case cmdAuthReq:
		resp, err = h.handleAuth(msg.Content)
	case cmdLoginReq:
		resp, err = h.handleLogin(msg.Content)
	}

	if err != nil {
		log.Error("handle err: ", err.Error())
	}
	return
}
func (h *handler)handleHeartbeat() {

}
func (h *handler)handleRegister(content string) (resp string, err error) {
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
	if err = h.dao.Call("User.Register", args, &reply); err != nil {
		return
	}
	log.Infof("User.Register %+v -> %+v", args, reply)
	respContent := &RegisterResp{reply.Code}
	resp = com.MakeMsgString(cmdRegisterResp, respContent)
	return
}

func (h *handler)handleAuth(content string) (resp string, err error) {
	var req AuthReq
	if err = json.Unmarshal([]byte(content), &req); err != nil {
		return
	}
	args := &dao.User_AuthArgs{req.Account, req.Psw}
	var reply dao.User_AuthReply
	if err = h.dao.Call("User.Auth", args, &reply); err != nil {
		return
	}
	respContent := &AuthResp{reply.Code}
	resp = com.MakeMsgString(cmdAuthResp, respContent)
	return
}

func (h *handler)handleLogin(content string) (resp string, err error) {
	resp = com.MakeMsgString(cmdLoginResp, "")
	return
}
package login
import (
	"encoding/json"
	"chess/com"
	"chess/dao"
	"chess/fw"
	log "github.com/lkj01010/log"
)

type handler struct {
	dc *dao.Client

}

func (h *handler)Handle(req string) (resp string, err error) {
	var msg com.Msg
	if err = json.Unmarshal([]byte(req), &msg); err != nil {
		log.Error("Unmarshal err: ", err)
		return
	}

	log.Debugf("msg: %#v", req)

	switch msg.Cmd {
	case cmdRegisterReq:
		resp, err = h.handleRegister(msg.Content)
	case cmdLoginReq:
		resp, err = h.handleLogin(msg.Content)
	}
	if err != nil{
		log.Error("handle err: ", err.Error())
	}
	return
}

func (h *handler)handleRegister(content string) (resp string, err error) {
	//	daocli.
	var req RegisterReq
	if err = json.Unmarshal([]byte(content), &req); err != nil {
		log.Error("content=", content, ", err: ", err.Error())
		return
	}
	var reply fw.RpcReply
	log.Debugf("req : %#v", req)
	if err = h.dc.UserRegister(req.Account, req.Psw, &reply); err != nil {
		return
	}
	resp = com.MakeMsgString(cmdRegisterResp, reply)
	return
}

func (h *handler)handleLogin(content string) (resp string, err error) {
	var req LoginReq
	if err = json.Unmarshal([]byte(content), &req); err != nil {
		return
	}
	var reply dao.UserAuthReply
	if err = h.dc.UserAuth(req.Account, req.Psw, &reply); err != nil {
		return
	}
	resp = com.MakeMsgString(cmdLoginResp, reply)
	return
}
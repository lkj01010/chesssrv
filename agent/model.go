package agent
import (
	"encoding/json"
	"chess/com"
	"chess/dao"
	log "github.com/lkj01010/log"
	"net/rpc"
	"chess/fw"
)

//数据处理模块
type model struct {
	dao     *rpc.Client

	agent   *fw.Agent

	id      string
	isLogin bool
}

//func NewModel(dao *rpc.Client, isTimeout bool) *model {
//	m := &model{
//		dao: dao,
//	}
//	return m
//}

func (m *model)Hook(a *fw.Agent) {
	m.agent = a
}

func (m *model)Enter() {
	m.isLogin = false
	m.id = ""
}

func (m *model)Exit() {
}

func (m *model)Handle(req string) (resp string, err error) {
	var msg com.Msg
	if err = json.Unmarshal([]byte(req), &msg); err != nil {
		return
	}

	if m.isLogin == false &&
	(msg.Cmd != com.Cmd_Com_Heartbeat &&
	msg.Cmd != com.Cmd_Ag_RegisterReq &&
	msg.Cmd != com.Cmd_Ag_AuthReq &&
	msg.Cmd != com.Cmd_Ag_LoginReq) {
		err = com.ErrCommandWithoutLogin
		return
	}

	switch msg.Cmd {
	case com.Cmd_Com_Heartbeat:
		m.handleHeartbeat()
	case com.Cmd_Ag_RegisterReq:
		resp, err = m.handleRegister(msg.Content)
	case com.Cmd_Ag_AuthReq:
		resp, err = m.handleAuth(msg.Content)
	case com.Cmd_Ag_LoginReq:
		resp, err = m.handleLogin(msg.Content)
	case com.Cmd_Ag_InfoReq:
		resp, err = m.handleInfo(msg.Content)

	case com.Cmd_Ag_ToGameReq:
		resp, err = m.handleToGame(msg.Content)
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
	var reply dao.Reply
	log.Debugf("req : %#v", req)
	//	if err = h.dc.UserRegister(&args, &reply); err != nil {
	if err = m.dao.Call("User.Register", args, &reply); err != nil {
		return
	}
	log.Infof("User.Register %+v -> %+v", args, reply)
	resp = com.MakeMsgString(com.Cmd_Ag_RegisterResp, reply.Code, nil)
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
	resp = com.MakeMsgString(com.Cmd_Ag_AuthResp, reply.Code, nil)
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
	if reply.Code == com.E_Success {
		//登录成功,记录用户id
		m.id = reply.Id
		m.isLogin = true
	}
	resp = com.MakeMsgString(com.Cmd_Ag_LoginResp, reply.Code, nil)
	return
}

func (m *model)handleInfo(content string) (resp string, err error) {
	args := &dao.Args{Id: m.id}
	log.Debugf("handleInfo args=%+v", args)
	var reply dao.User_InfoReply
	if err = m.dao.Call("User.GetInfo", args, &reply); err != nil {
		return
	}
	log.Debugf("handleInfo, reply=%+v", reply)
	resp = com.MakeMsgString(com.Cmd_Ag_InfoResp, reply.Code, reply.Info)
	return
}

func (m *model)handleToGame(content string) (resp string, err error) {
	req := &ToRoomReq{Id: m.id, Content: content}
	var b []byte
	if b, err = json.Marshal(req); err != nil {
		log.Error("handleToGame:error=", err)
	}
	msg := com.MakeConnIdRawMsgString(m.agent.ConnId, b)    // can? b is of []type, need string
	serverInst.gameCli.Send(msg)
	return
}

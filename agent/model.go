package agent
import (
	"encoding/json"
	"chess/com"
	"chess/dao"
	log "github.com/lkj01010/log"
	"net/rpc"
	"chess/fw"
	"errors"
	"fmt"
)

//数据处理模块
type model struct {
	dao     *rpc.Client

	agent   *fw.Agent

	// 登录成功后赋值.通过是否为""判断是否登录
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
	if m.isLogin {
		serverInst.RemoveFromLoginAgent(m.id)
	}
}

func (m *model)Handle(req string) (resp string, err error) {
	var msg com.Msg
	if err = json.Unmarshal([]byte(req), &msg); err != nil {
//		log.Error("Unmarshal err: ", err)
		return
	}

	if m.isLogin == false &&
	(msg.Cmd != CmdHeartbeat &&
	msg.Cmd != CmdRegisterReq &&
	msg.Cmd != CmdAuthReq &&
	msg.Cmd != CmdLoginReq) {
		e := fmt.Sprintf("Handle:not allowed withou LOGIN:cmd=%+v", msg.Cmd)
		err = errors.New(e)
		return
	}

	switch msg.Cmd {
	case CmdHeartbeat:
		m.handleHeartbeat()
	case CmdRegisterReq:
		resp, err = m.handleRegister(msg.Content)
	case CmdAuthReq:
		resp, err = m.handleAuth(msg.Content)
	case CmdLoginReq:
		resp, err = m.handleLogin(msg.Content)
	case CmdInfoReq:
		resp, err = m.handleInfo(msg.Content)

	case Cmd_Ag_ToGameReq:
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
	var reply dao.RpcReply
	log.Debugf("req : %#v", req)
	//	if err = h.dc.UserRegister(&args, &reply); err != nil {
	if err = m.dao.Call("User.Register", args, &reply); err != nil {
		return
	}
	log.Infof("User.Register %+v -> %+v", args, reply)
	resp = com.MakeMsgString(CmdRegisterResp, reply.Code, nil)
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
	resp = com.MakeMsgString(CmdAuthResp, reply.Code, nil)
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
	resp = com.MakeMsgString(CmdLoginResp, reply.Code, nil)
	return
}

func (m *model)handleInfo(content string) (resp string, err error) {
	args := &dao.User_InfoArgs{Id: m.id}
	log.Debugf("handleInfo args=%+v", args)
	var reply dao.User_InfoReply
	if err = m.dao.Call("User.Info", args, &reply); err != nil {
		return
	}
	log.Debugf("handleInfo, reply=%+v", reply)
	resp = com.MakeMsgString(CmdInfoResp, reply.Code, reply.Info)
	return
}

func (m *model)handleToGame(content string) (resp string, err error) {
	req := &ToRoomReq{Id: m.id, Content: content}
	var b []byte
	if b, err = json.Marshal(req); err != nil {
		log.Error("handleToGame:error=", err)
	}
	msg := com.MakeConnIdRawMsgString(m.agent.ConnId, b)	// can? b is of []type, need string
	serverInst.gameCli.Send(msg)
	return
}

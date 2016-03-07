package agent
//import (
//	"github.com/lkj01010/log"
//	"time"
//	"chess/com"
//	"encoding/json"
//	"chess/dao"
//	"net/rpc"
//)
//
//const (
//	CtrlRemoveAgent = "removeAgent"
//)
//
//type ReadWriteCloser interface {
//	Read(msg *string) error
//	Write(msg string) error
//	Close() error
//}
//
//type agent struct {
//	ReadWriteCloser
//	ctrl chan string
//	dao  *rpc.Client
//
//	// 登录成功后赋值.通过是否为""判断是否登录
//	id   string
//}
//
//func NewAgent(rwc ReadWriteCloser, dao *rpc.Client) *agent {
//	return &agent{
//		ReadWriteCloser: rwc,
//		ctrl: make(chan string, 10),
//		dao: dao,
//		id: "",
//	}
//}
//
//func (a *agent)Serve() (err error) {
//	a.enter()
//	session := make(chan string, 1)
//	go func(c chan string) {
//		var buf string
//		for {
//			err = a.Read(&buf)
//			if err != nil {
//				//fixme: 通过正常流程退出
//				log.Debug("agent read: ", err.Error())
//				return
//			}
//			c <- buf
//		}
//	}(session)
//
//	timeout := time.NewTimer(0)
//	L:    for {
//		timeout.Reset(10 * time.Second)
//		select {
//		case ctrl := <-a.ctrl:
//			if ctrl == CtrlRemoveAgent {
//				log.Debug("recv ctrl: CtrlRemoveAgent")
//				break L
//			}
//
//		case msg := <-session:
//			var resp string
//			resp, err = a.handle(msg)
//			if err != nil {
//				log.Error("agent session: ", err.Error())
//				return
//			}
//			a.Write(resp)
//
//		case <-timeout.C:
//			log.Debug("recv timeout")
//			break L
//		}
//	}
//
//	a.exit()
//
//	return
//}
//
//func (a *agent)enter() {
//
//}
//
//func (a *agent)exit() {
//	if a.id != "" {
//		serverInst.RemoveAgent(a.id)
//	}
//}
//
//func (a *agent)handle(req string) (resp string, err error) {
//	var msg com.Msg
//	if err = json.Unmarshal([]byte(req), &msg); err != nil {
//		log.Error("Unmarshal err: ", err)
//		return
//	}
//
//	switch msg.Cmd {
//	case CmdHeartbeat:
//		a.handleHeartbeat()
//	case CmdRegisterReq:
//		resp, err = a.handleRegister(msg.Content)
//	case CmdAuthReq:
//		resp, err = a.handleAuth(msg.Content)
//	case CmdLoginReq:
//		resp, err = a.handleLogin(msg.Content)
//	case CmdInfoReq:
//		resp, err = a.handleInfo(msg.Content)
//	}
//	if err != nil {
//		log.Error("handle err: ", err.Error())
//	}
//	return
//}
//func (a *agent)handleHeartbeat() {
//	// do nothing
//}
//
//func (a *agent)handleRegister(content string) (resp string, err error) {
//	//	daocli.
//	var req RegisterReq
//	if err = json.Unmarshal([]byte(content), &req); err != nil {
//		log.Error("content=", content, ", err: ", err.Error())
//		return
//	}
//	args := &dao.User_RegisterArgs{req.Account, req.Psw}
//	var reply dao.RpcReply
//	log.Debugf("req : %#v", req)
//	//	if err = h.dc.UserRegister(&args, &reply); err != nil {
//	if err = a.dao.Call("User.Register", args, &reply); err != nil {
//		return
//	}
//	log.Infof("User.Register %+v -> %+v", args, reply)
//	resp = com.MakeMsgString(CmdRegisterResp, reply.Code, nil)
//	return
//}
//
//func (a *agent)handleAuth(content string) (resp string, err error) {
//	var req AuthReq
//	if err = json.Unmarshal([]byte(content), &req); err != nil {
//		return
//	}
//	args := &dao.User_AuthArgs{req.Account, req.Psw}
//	var reply dao.User_AuthReply
//	if err = a.dao.Call("User.Auth", args, &reply); err != nil {
//		return
//	}
//	resp = com.MakeMsgString(CmdAuthResp, reply.Code, nil)
//	return
//}
//
//func (a *agent)handleLogin(content string) (resp string, err error) {
//	var req LoginReq
//	if err = json.Unmarshal([]byte(content), &req); err != nil {
//		return
//	}
//	log.Infof("handleLogin, req=%+v", req)
//	args := &dao.User_AuthArgs{req.Account, req.Psw}
//	var reply dao.User_AuthReply
//	if err = a.dao.Call("User.Auth", args, &reply); err != nil {
//		return
//	}
//	log.Infof("handleLogin, reply=%+v", reply)
//	if reply.Code == com.E_Success {
//		//登录成功,记录用户id
//		a.id = reply.Id
//		// 记录到server
//		serverInst.AddAgent(a.id, a)
//	}
//	resp = com.MakeMsgString(CmdLoginResp, reply.Code, nil)
//	return
//}
//
//func (a *agent)handleInfo(content string) (resp string, err error) {
//	args := &dao.User_InfoArgs{Id: a.id}
//	log.Debugf("handleInfo args=%+v", args)
//	var reply dao.User_InfoReply
//	if err = a.dao.Call("User.Info", args, &reply); err != nil {
//		return
//	}
//	log.Debugf("handleInfo, reply=%+v", reply)
//	resp = com.MakeMsgString(CmdInfoResp, reply.Code, reply.Info)
//	return
//}

package agent
import (
	"github.com/lkj01010/log"
	"time"
	"chess/com"
	"encoding/json"
	"chess/dao"
	"net/rpc"
)

const (
	CtrlRemoveAgent = "removeAgent"
)

type ReadWriteCloser interface {
	Read(msg *string) error
	Write(msg string) error
	Close() error
}

type agent struct {
	ReadWriteCloser
	ctrl chan string

	dao  *rpc.Client
	// 登录成功后赋值.通过是否为""判断是否登录
	id   string
}

func NewAgent(rwc ReadWriteCloser, dao *rpc.Client) *agent {
	return &agent{
		ReadWriteCloser: rwc,
		ctrl: make(chan string, 0),
		dao: dao,
		id: "",
	}
}

func (a *agent)Serve() (err error) {
	a.enter()
	session := make(chan string, 0)
	go func(c chan string) {
		var buf string
		for {
			err = a.Read(&buf)
			if err != nil {
				//fixme: 通过正常流程退出
				log.Debug("agent read: ", err.Error())
				return
			}
			c <- buf
		}
	}(session)

	timeout := time.NewTimer(0)
	L:    for {
		timeout.Reset(10 * time.Second)
		select {
		case msg := <-session:
			log.Debug("recv msg")
			var resp string
			resp, err = a.handle(msg)
			if err != nil {
				log.Error("agent session: ", err.Error())
				return
			}
			a.Write(resp)

		case <-timeout.C:
			log.Debug("recv timeout")
			break L
		case ctrl := <-a.ctrl:
			if ctrl == CtrlRemoveAgent {
				log.Debug("recv ctrl: CtrlRemoveAgent")
				break L
			}
		}
	}

	a.exit()

	return
}

func (a *agent)enter() {

}

func (a *agent)exit() {

}

func (a *agent)handle(req string) (resp string, err error) {
	var msg com.Msg
	if err = json.Unmarshal([]byte(req), &msg); err != nil {
		log.Error("Unmarshal err: ", err)
		return
	}

	switch msg.Cmd {
	case cmdHeartbeat:
		a.handleHeartbeat()
	case cmdRegisterReq:
		resp, err = a.handleRegister(msg.Content)
	case cmdAuthReq:
		resp, err = a.handleAuth(msg.Content)
	case cmdLoginReq:
		resp, err = a.handleLogin(msg.Content)
	case cmdInfoReq:
		resp, err = a.handleInfo(msg.Content)
	}
	if err != nil {
		log.Error("handle err: ", err.Error())
	}
	return
}
func (a *agent)handleHeartbeat() {
	// do nothing
}

func (a *agent)handleRegister(content string) (resp string, err error) {
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
	if err = a.dao.Call("User.Register", args, &reply); err != nil {
		return
	}
	log.Infof("User.Register %+v -> %+v", args, reply)
	resp = com.MakeMsgString(cmdRegisterResp, reply.Code, nil)
	return
}

func (a *agent)handleAuth(content string) (resp string, err error) {
	var req AuthReq
	if err = json.Unmarshal([]byte(content), &req); err != nil {
		return
	}
	args := &dao.User_AuthArgs{req.Account, req.Psw}
	var reply dao.User_AuthReply
	if err = a.dao.Call("User.Auth", args, &reply); err != nil {
		return
	}
	resp = com.MakeMsgString(cmdAuthResp, reply.Code, nil)
	return
}

func (a *agent)handleLogin(content string) (resp string, err error) {
	var req LoginReq
	if err = json.Unmarshal([]byte(content), &req); err != nil {
		return
	}
	args := &dao.User_AuthArgs{req.Account, req.Psw}
	var reply dao.User_AuthReply
	if err = a.dao.Call("User.Auth", args, &reply); err != nil {
		return
	}
	if reply.Code == com.E_Success {
		//登录成功,记录用户id
		a.id = reply.Id
	}
	resp = com.MakeMsgString(cmdLoginResp, reply.Code, nil)
	return
}

func (a *agent)handleInfo(content string) (resp string, err error) {
	args := &dao.User_InfoArgs{Id: a.id}
	log.Debugf("handleInfo args=%+v", args)
	var reply dao.User_InfoReply
	if err = a.dao.Call("User.Info", args, &reply); err != nil {
		return
	}
	log.Debugf("handleInfo, reply=%+v", reply)
	resp = com.MakeMsgString(cmdInfoResp, reply.Code, reply.Info)
	return
}

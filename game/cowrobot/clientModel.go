package cowrobot
import (
//	"github.com/lkj01010/log"
//	"golang.org/x/net/websocket"
//	"time"
//	"chess/cfg"
//	"net"
	"github.com/lkj01010/log"
	"chess/com"
	"encoding/json"
	"chess/agent"
)


type Clientmodel struct {
}

func (c *Clientmodel)Handle(req string) (err error) {
	var msg com.Msg
	if err = json.Unmarshal([]byte(req), &msg); err != nil {
		log.Error("Unmarshal err: ", err)
		return
	}

	switch msg.Cmd {
	case agent.CmdRegisterResp:
		err = c.handleRegister(msg.Content)
	case agent.CmdAuthResp:
		err = c.handleAuth(msg.Content)
	case agent.CmdLoginResp:
		err = c.handleLogin(msg.Content)
	case agent.CmdInfoResp:
		err = c.handleInfo(msg.Content)
	}
	if err != nil {
		log.Error("handle err: ", err.Error())
	}
	return
}

func (c *Clientmodel)handleRegister(content string) (err error) {
	log.Info("register")
	return
}

func (c *Clientmodel)handleAuth(content string) (err error) {
	log.Info("handleAuth")
	return
}

func (c *Clientmodel)handleLogin(content string) (err error) {
	log.Info("handleLogin")
	return
}

func (c *Clientmodel)handleInfo(content string) (err error) {
	log.Info("handleInfo")
	return
}

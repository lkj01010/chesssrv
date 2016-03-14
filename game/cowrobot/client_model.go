package cowrobot
import (
	"github.com/lkj01010/log"
	"chess/com"
	"encoding/json"
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
	case com.Cmd_Ag_RegisterResp:
		err = c.handleRegister(msg.Content)
	case com.Cmd_Ag_AuthResp:
		err = c.handleAuth(msg.Content)
	case com.Cmd_Ag_LoginResp:
		err = c.handleLogin(msg.Content)
	case com.Cmd_Ag_InfoResp:
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

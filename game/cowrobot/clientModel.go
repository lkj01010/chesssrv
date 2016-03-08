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
	"time"
)


type clientmodel struct {
}

func (c *clientmodel)Step(output *string){
	*output = string(`{"cmd":104,"content":"{\"account\":\"testUtf\",\"psw\":\"pswlk22\"}"}`)
	log.Debugf("Step: send=%+v", output)

//	time.Sleep(1*time.Second)
	time.Sleep(1)
}

func (c *clientmodel)Handle(req string) (err error) {
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

func (c *clientmodel)handleRegister(content string) (err error) {
	log.Info("register")
	return
}

func (c *clientmodel)handleAuth(content string) (err error) {
	log.Info("handleAuth")
	return
}

func (c *clientmodel)handleLogin(content string) (err error) {
	log.Info("handleLogin")
	return
}

func (c *clientmodel)handleInfo(content string) (err error) {
	log.Info("handleInfo")
	return
}

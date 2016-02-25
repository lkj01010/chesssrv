package login
import (
"encoding/json"
"chess/com"
	"chess/dao"
	"chess/fw"
"github.com/lkj01010/log"
)

type handler struct {
	dc *dao.Client
}

func (h *handler)Handle(req string) (resp string, err error) {
	//	err := json.Unmarshal(req, )
	var msg com.Msg
	if err = json.Unmarshal([]byte(req), &msg); err != nil {
		return
	}

	switch msg.Cmd {
	case cmdRegisterReq:
		resp = h.handleRegister(msg.Content)
	}
	return
}

func (h *handler)handleRegister(content string) (resp string, err error){
	//	daocli.
	var req RegisterReq
	if err = json.Unmarshal([]byte(content), &req); err != nil {
		return
	}
	var reply fw.RpcReply
	if err = h.dc.UserRegister(req.Account, req.Psw, &reply); err != nil {
		return
	}
}
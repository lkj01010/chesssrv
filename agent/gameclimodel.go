package agent

import (
	"github.com/lkj01010/log"
	"chess/com"
	"encoding/json"
)

type gameCliModel struct {
}

func (m *gameCliModel)Handle(req string) (err error) {
	var msg com.ConnIdRawMsg
	if err = json.Unmarshal([]byte(req), &msg); err != nil {
		log.Error("Unmarshal err: ", err)
		return
	}

	agent := serverInst.GetAgent(msg.ConnId)
	if agent != nil {
		agent.Send(msg.Content)
	} else {
		err = com.ErrAgentNotFound
	}
	return
}

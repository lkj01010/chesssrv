package agent

import (
	"github.com/lkj01010/log"
	"chess/com"
	"encoding/json"
)


type gameCliModel struct {
}

func (c *gameCliModel)Handle(req string) (err error) {
	var msg com.Msg
	if err = json.Unmarshal([]byte(req), &msg); err != nil {
		log.Error("Unmarshal err: ", err)
		return
	}

	if err != nil {
		log.Error("handle err: ", err.Error())
	}
	return
}

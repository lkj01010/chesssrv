package game
import (
	"chess/com"
	"encoding/json"
)


const (
	Cmd_Game_Start = 3000
	Cmd_Game_EnterReq = 3000
	Cmd_Game_EnterResp = 3001

	Cmd_Game_LeaveReq = 3010
	Cmd_Game_LeaveNtf = 3012

	Cmd_Game_End = 3499
)

type Msg struct {
	Id  string `json:"id"`
	Msg com.Msg `json:"msg"`
}

func MakeMsgString(id string, cmd int, param int, content interface{}) (resp string) {
	var msg Msg
	msg.Id = id
	msg.Msg = com.MakeMsgString(cmd, param, content)
	r, _ := json.Marshal(msg)
	return string(r)
}

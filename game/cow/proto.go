package cow
import (
	"encoding/json"
	"chess/com"
)

const (
	Cmd_Cow_Start = 3500
	Cmd_Cow_DealNtf = 3502
	Cmd_Cow_CardTypeShowReq = 3510
	Cmd_Cow_CardTypeShowNtf = 3511
	Cmd_Cow_SettleNtf = 3522

	Cmd_Cow_NotEnoughMoneyNtf = 3702

	Cmd_Cow_End = 3999
)

type Msg struct {
	Id  string `json:"id"`
//	Msg com.Msg `json:"msg"`
	Content string `json:"content"`
}

func MakeMsgString(id string, cmd int, param int, content interface{}) (resp string) {
	var msg Msg
	msg.Id = id
	msg.Content = com.MakeMsgString(cmd, param, content)
	r, _ := json.Marshal(msg)
	return string(r)
}


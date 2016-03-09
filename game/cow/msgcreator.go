package cow
import (
	"chess/com"
)

type msgCreator struct{}
var msgcInst msgCreator

func init() {
	msgcInst = msgCreator{}
}

func (mc *msgCreator)hasNoEnoughMoney(id int) string {
	content := com.MakeMsgString(Cmd_Cow_NotEnoughMoneyNtf, 0, nil)
	msg := com.MakeConnIdRawMsgString(id, content)
	return msg
}



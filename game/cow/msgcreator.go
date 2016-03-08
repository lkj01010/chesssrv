package cow
import (
)

type msgCreator struct{}
var msgcInst msgCreator

func init() {
	msgcInst = msgCreator{}
}

func (mc *msgCreator)hasNoEnoughMoney(id string) string {
	msg := MakeMsgString(id, Cmd_Cow_NotEnoughMoneyNtf, 0, nil)
	return msg
}



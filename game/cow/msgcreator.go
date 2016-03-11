package cow
import (
	"chess/com"
)

type msgCreator struct{}
var msgCreatorInst msgCreator

func init() {
	msgCreatorInst = msgCreator{}
}

func (mc *msgCreator)hasNoEnoughMoney() string {
	return com.MakeMsgString(Cmd_Cow_NotEnoughMoneyNtf, 0, nil)
}

func (mc *msgCreator)Settle(settleNtf *SettleNtf) string {
	return com.MakeMsgString(Cmd_Cow_SettleNtf, 0, settleNtf)
}



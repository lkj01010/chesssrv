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
	return com.MakeMsgString(com.Cmd_Cow_NotEnoughMoneyNtf, 0, nil)
}

func (mc *msgCreator)Settle(settleNtf *SettleNtf) string {
	return com.MakeMsgString(com.Cmd_Cow_SettleNtf, 0, settleNtf)
}

func (mc *msgCreator)PlayerEnter(idx int, info *com.UserInfo) string {
	ntf := &PlayerEnterNtf{
		Info: *info,
		Index: idx,
	}
	return com.MakeMsgString(com.Cmd_Cow_PlayerEnterNtf, 0, ntf)
}

func (mc *msgCreator)WaitReady() string {
	//todo: 0314
//	return com.MakeMsgString(Cmd_Cow_PlayerEnterNtf)
	return ""
}
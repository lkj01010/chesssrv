package cow
import (
	"chess/game"
)

type msgCreator int
var msgcInst msgCreator

func init() {
	msgcInst = msgCreator{}
}

func (mc *msgCreator)hasNoEnoughMoney(id string) string {
	msg := game.MakeMsgString(id, Cmd_Cow_NotEnoughMoneyNtf, 0, nil)
	return msg
}



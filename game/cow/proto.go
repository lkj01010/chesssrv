package cow
import (
	"chess/com"
)

const (
	Cmd_Cow_PlayerEnterNtf = 3500
	Cmd_Cow_WaitReadyNtf = 3501
	Cmd_Cow_Ready = 3502
	Cmd_Cow_DealNtf = 3503
	Cmd_Cow_CardTypeShowReq = 3510
	Cmd_Cow_CardTypeShowNtf = 3511
	Cmd_Cow_SettleNtf = 3522

	Cmd_Cow_NotEnoughMoneyNtf = 3702

	Cmd_Cow_End = 3999
)

const (
	Cmd_Inner_PlayerEnter = 4000
)

type PlayerEnterNtf struct {
	Info com.UserInfo `json:"info"`
	Index int `json:"index"`
}

type SettleNtf struct {
	TempString string `json:"tempstring"`
}
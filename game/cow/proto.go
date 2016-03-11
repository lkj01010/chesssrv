package cow
import (
)

const (
	Cmd_Cow_Ready = 3500
	Cmd_Cow_DealNtf = 3502
	Cmd_Cow_CardTypeShowReq = 3510
	Cmd_Cow_CardTypeShowNtf = 3511
	Cmd_Cow_SettleNtf = 3522

	Cmd_Cow_NotEnoughMoneyNtf = 3702

	Cmd_Cow_End = 3999
)

type SettleNtf struct {
	TempString string `json:"tempstring"`
}


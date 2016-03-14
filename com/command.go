package com

type Command int

// com
const (
	Cmd_Com_Start Command = iota +500
	Cmd_Com_Heartbeat
	Cmd_Com_End
)

// agent
const (
	Cmd_Ag_Start Command = iota + 1000
	Cmd_Ag_RegisterReq
	Cmd_Ag_RegisterResp
	Cmd_Ag_AuthReq
	Cmd_Ag_AuthResp
	Cmd_Ag_LoginReq
	Cmd_Ag_LoginResp
	Cmd_Ag_InfoReq
	Cmd_Ag_InfoResp

	Cmd_Ag_ToGameReq
	Cmd_Ag_End
)

// game
const (
	Cmd_Game_Start Command = iota + 2000
	Cmd_Game_EnterReq
	Cmd_Game_EnterResp

	Cmd_Game_LeaveReq
	Cmd_Game_LeaveNtf

	Cmd_Game_End
)

// cow
const (
	Cmd_Cow_Start Command = iota + 3000
	Cmd_Cow_PlayerEnterNtf
	Cmd_Cow_WaitReadyNtf
	Cmd_Cow_Ready
	Cmd_Cow_DealNtf
	Cmd_Cow_CardTypeShowReq
	Cmd_Cow_CardTypeShowNtf
	Cmd_Cow_SettleNtf

	Cmd_Cow_NotEnoughMoneyNtf

	Cmd_Cow_End
)




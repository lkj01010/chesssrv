package com

type Cmd int
const (
	Cmd_Game_Start Cmd = iota
	Cmd_Game_EnterReq
	Cmd_Game_EnterResp

	Cmd_Game_LeaveReq
	Cmd_Game_LeaveNtf

	Cmd_Game_End
)

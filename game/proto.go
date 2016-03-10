package game
import (
)

const (
	Cmd_Game_Start = 3000
	Cmd_Game_EnterReq = 3000
	Cmd_Game_EnterResp = 3001

	Cmd_Game_LeaveReq = 3010
	Cmd_Game_LeaveNtf = 3012

	Cmd_Game_End = 3499
)

type EnterGame struct {
	Id string `json:"id"`
	RoomType int `json:"roomtype"`
}
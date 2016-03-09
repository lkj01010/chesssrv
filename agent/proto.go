package agent
import (
)

const (
	CmdHeartbeat = 10
	CmdRegisterReq = 100
	CmdRegisterResp = 101

	CmdAuthReq = 102
	CmdAuthResp = 103
	CmdLoginReq = 104
	CmdLoginResp = 105
	CmdInfoReq = 106
	CmdInfoResp = 107

	CmdToLobbyReq = 200
	CmdToLobbyResp = 201

	Cmd_Ag_ToGameReq = 202
)

var cmdText = &map[int]string{
	CmdHeartbeat:	"heartbeat",
	CmdRegisterReq:    "registerReq",
	CmdRegisterResp:    "registerResp",
	CmdAuthReq:    "authReq",
	CmdAuthResp:    "authResp",
	CmdLoginReq:    "loginReq",
	CmdLoginResp:    "loginResp",

	CmdToLobbyReq:    "toLobbyReq",
	CmdToLobbyResp:    "toLobbyResp",
	Cmd_Ag_ToGameReq:    "toGameReq",
}

//var loginMethodCodes map[string]int
//
//func init() {
//	if loginMethodCodes, err := json.Marshal(loginMethodCodeJson); err != nil {
//		fw.Log.WithField("json.Marshal.err", err.Error()).Error("login:proto:init error")
//	}
//}

//用户认证
type AuthReq struct {
	Account string `json:"account"`
	Psw     string `json:"psw"`
}

type RegisterReq struct {
	Account string `json:"account"`
	Psw     string `json:"psw"`
}

type LoginReq struct {
	Account string `json:"account"`
	Psw     string `json:"psw"`
}


/////////////////////////////////////////////
// remote

type ToLobbyReq struct {
	Msg string `json:"msg"`
}

type FromLobbyResp struct {
	Msg string `json:"msg"`
}

type ToRoomReq struct {
	Id string `json:"id"`
	Content string `json:"content"`
}

type FromRoomResp struct {
	Id string `json:"id"`
	Cmd string `json:"cmd"`
	Msg string `json:"msg"`
}

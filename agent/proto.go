package agent
import (
)

const (
	cmdHeartbeat = 10
	cmdRegisterReq = 100
	cmdRegisterResp = 101

	cmdAuthReq = 102
	cmdAuthResp = 103
	cmdLoginReq = 104
	cmdLoginResp = 105

	cmdToLobbyReq = 200
	cmdToLobbyResp = 201

	cmdToGameReq = 202
	cmdToGameResp = 203
)

var cmdText = &map[int]string{
	cmdHeartbeat:	"heartbeat",
	cmdRegisterReq:    "registerReq",
	cmdRegisterResp:    "registerResp",
	cmdAuthReq:    "authReq",
	cmdAuthResp:    "authResp",
	cmdLoginReq:    "loginReq",
	cmdLoginResp:    "loginResp",

	cmdToLobbyReq:    "toLobbyReq",
	cmdToLobbyResp:    "toLobbyResp",
	cmdToGameReq:    "toGameReq",
	cmdToGameResp:    "toGameResp",
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

type ToLobbyReq struct {
	Msg string `json:"msg"`
}

type ToLobbyResp struct {
	Msg string `json:"msg"`
}

type ToRoomReq struct {
	Msg string `json:"msg"`
}

type ToRoomResp struct {
	Msg string `json:"msg"`
}

package login
import (
	"encoding/json"
)

/*
	proto should be:
	{
		method: int,
		params: obj
	}

 */

const (
	methodRegisterReq = 100
	methodRegisterResp = 101

	methodLoginReq = 102
	methodLoginResp = 103

	methodToLobbyReq = 200
	methodToLobbyResp = 201

	methodToGameReq = 202
	methodToGameResp = 203
)

const methodText = &map[int]string{
	methodRegisterReq:    "registerReq",
	methodRegisterResp:    "registerResp",
	methodLoginReq:    "loginReq",
	methodLoginResp:    "loginResp",
	methodToLobbyReq:    "toLobbyReq",
	methodToLobbyResp:    "toLobbyResp",
	methodToGameReq:    "toGameReq",
	methodToGameResp:    "toGameResp",
}

//var loginMethodCodes map[string]int
//
//func init() {
//	if loginMethodCodes, err := json.Marshal(loginMethodCodeJson); err != nil {
//		fw.Log.WithField("json.Marshal.err", err.Error()).Error("login:proto:init error")
//	}
//}

type LoginReq struct {
	Account string `json:"account"`
	Psw     string `json:"psw"`
}

type LoginResp struct {
	Code int `json:"code"`
}

type RegisterReq struct {
	Account string `json:"account"`
	Psw     string `json:"psw"`
}

type RegisterResp struct {
	Code int `json:"code"`
}

type ToLobbyReq struct {
	Msg string `json:"msg"`
}

type ToLobbyResp struct {
	Msg	string `json:"msg"`
}

type ToGameReq struct {
	Msg string `json:"msg"`
}

type ToGameResp struct {
	Msg string `json:"msg"`
}

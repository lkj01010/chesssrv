package login
import (
)

/*
	proto should be:
	{
		method: int,
		params: obj
	}

 */

const (
	cmdRegisterReq = 100
	cmdRegisterResp = 101

	cmdLoginReq = 102
	cmdLoginResp = 103

	cmdToLobbyReq = 200
	cmdToLobbyResp = 201

	cmdToGameReq = 202
	cmdToGameResp = 203
)

var cmdText = &map[int]string{
	cmdRegisterReq:    "registerReq",
	cmdRegisterResp:    "registerResp",
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

type ToRoomReq struct {
	Msg string `json:"msg"`
}

type ToRoomResp struct {
	Msg string `json:"msg"`
}

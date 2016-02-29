package agent
import (
)


const (
	cmdLoginReq = 100
	cmdLoginResp = 101


	cmdToLobbyReq = 200
	cmdToLobbyResp = 201

	cmdToGameReq = 202
	cmdToGameResp = 203
)

var cmdText = &map[int]string{
	cmdLoginReq:    "loginReq",
	cmdLoginResp:    "loginResp",
	cmdToLobbyReq:    "toLobbyReq",
	cmdToLobbyResp:    "toLobbyResp",
	cmdToGameReq:    "toGameReq",
	cmdToGameResp:    "toGameResp",
}

type LoginReq struct {
	Account string `json:"account"`
	Psw     string `json:"psw"`
}

type LoginResp struct {
	Code int `json:"code"`
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

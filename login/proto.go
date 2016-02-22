package login
import (
)

const (
	methodRegister = 1000
	methodLogin = 1001
)

const methodText = &map[int]string{
	methodRegister:    "register",
	methodLogin:    "login",
}

//var loginMethodCodes map[string]int
//
//func init() {
//	if loginMethodCodes, err := json.Marshal(loginMethodCodeJson); err != nil {
//		fw.Log.WithField("json.Marshal.err", err.Error()).Error("login:proto:init error")
//	}
//}

type ReqLogin struct {
	Account string `json:"account"`
	Psw     string `json:"psw"`
}

type RespLogin struct {
	Code int    `json:"code"`
}

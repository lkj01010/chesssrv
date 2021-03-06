package com
import "encoding/json"

///////////////////////////////////////////////////////
type Msg struct {
	Cmd     int `json:"cmd"`
	Param   int `json:"param"`
	Content string `json:"content"`
}

func MakeMsgString(cmd Command, param int, content interface{}) string {
	var msg Msg
	msg.Cmd = int(cmd)
	msg.Param = param

	if content != nil {
		c, _ := json.Marshal(content)
		msg.Content = string(c)
	}
	r, _ := json.Marshal(msg)
	return string(r)
}

///////////////////////////////////////////////////////
// deprecated
//	Id  string `json:"id"`
//	Cmd int `json:"cmd"`
//	Param int `json:"param"`
//	Content string `json:"content"`
//}
//
//func MakeIdMsgString(id string, cmd Command, param int, content interface{}) string {
//	var msg IdMsg
//	msg.Id = id
//	msg.Cmd = int(cmd)
//	msg.Param = param
//	if content != nil {
//		c, _ := json.Marshal(content)
//		msg.Content = string(c)
//	}
//	r, _ := json.Marshal(msg)
//	return string(r)
//}

///////////////////////////////////////////////////////
type ConnIdRawMsg struct {
	ConnId  int `json:"connid"`
	Content string `json:"content"`
}

func MakeConnIdRawMsgString(connId int, content interface{}) string {
	var msg ConnIdRawMsg
	msg.ConnId = connId
	if content != nil {
		c, _ := json.Marshal(content)
		msg.Content = string(c)
	}
	r, _ := json.Marshal(msg)
	return string(r)
}

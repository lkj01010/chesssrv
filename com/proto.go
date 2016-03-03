package com
import "encoding/json"

type Msg struct {
	Cmd     int `json:"cmd"`
	Code	int `json:"code"`
	Content string `json:"content"`
}

func MakeMsgString(cmd int, code int, content interface{}) (resp string) {
	var msg Msg
	msg.Cmd = cmd
	msg.Code = code

	if content != nil {
		c, _ := json.Marshal(content)
		msg.Content = string(c)
	}
	r, _ := json.Marshal(msg)
	return string(r)
}

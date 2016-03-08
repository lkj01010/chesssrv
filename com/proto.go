package com
import "encoding/json"

type Msg struct {
	Cmd     int `json:"cmd"`
	Param   int `json:"param"`
	Content string `json:"content"`
}

func MakeMsgString(cmd int, param int, content interface{}) (resp string) {
	var msg Msg
	msg.Cmd = cmd
	msg.Param = param

	if content != nil {
		c, _ := json.Marshal(content)
		msg.Content = string(c)
	}
	r, _ := json.Marshal(msg)
	return string(r)
}

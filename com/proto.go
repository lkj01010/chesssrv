package com
import "encoding/json"

type Msg struct {
	Cmd     int `json:"cmd"`
	Content string `json:"content"`
}

func MakeMsgString(cmd int, content interface{}) (resp string) {
	var msg Msg
	msg.Cmd = cmd

	b, _ := json.Marshal(content)
	msg.Content = string(b)
	b, _ = json.Marshal(msg)
	return string(b)
}

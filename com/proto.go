package com
import "encoding/json"

type Msg struct {
	Cmd int `json:"cmd"`
	Content string `json:"cotent"`
}

func MakeMsgString(cmd int, content interface{})(resp string){
	var msg Msg
	msg.Cmd = cmd
	msg.Content = json.Marshal(content)
	return json.Marshal(msg)
}
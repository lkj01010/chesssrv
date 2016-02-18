package data

import (
	"testing"
	"chess/fw"
)

func TestLogin(t *testing.T){
	u := new(User)
	u.Init()

	//
	code, _ := u.HandleRegister("lkj", "lkjpassword")
	fw.Log.WithField("code", code).Info("TestHandleRegister")

	u.Exit()
}

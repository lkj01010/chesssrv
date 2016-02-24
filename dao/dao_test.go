package dao

import (
	"testing"
)

func TestLogin(t *testing.T){
	defer func(){
		Exit()
	}()

	//
	var registerReply int
	user.HandleRegister(&RegisterArgs{Account:"lkj2", Psw:"lkjpassword"}, &registerReply)

	var reply AuthArgs
	user.HandleAuth(&AuthArgs{Account:"lkj2", Psw:"lkjpassword"}, &reply)
}

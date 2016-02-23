package data

import (
	"testing"
)

func TestLogin(t *testing.T){
	u := new(User)
	u.Init()
	g := new(Game)
	g.Init()
	defer func(){
		u.Exit()
		g.Exit()
	}()

	//
	u.HandleRegister("lkj2", "lkjpassword")

	u.HandleAuth(g, "lkj2", "lkjpassword")

	u.Exit()
}

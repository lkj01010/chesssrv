package data
import (
	"github.com/garyburd/redigo/redis"
	"chess/fw"
	"github.com/Sirupsen/logrus"
	"chess/com"
	"chess/cfg"
	"time"
)

type User struct {
	c redis.Conn
}

func(u *User)prepare()(b bool, err error){
	//setup connection
	u.c, err = redis.Dial("tcp", cfg.RedisAddr(),
		redis.DialReadTimeout(1 * time.Second), redis.DialWriteTimeout(1 * time.Second))
	if err != nil {
		fw.Log.Error("redis.Dial error")
		return
	}

	//fill keys
	b, _ = u.c.Do("EXIST", "account:count")
	if b == false {
		u.c.Do("SET", "account:count", 0)
	}
	return
}

func(u *User)HandleRegister(account, psw string)(code int, err error){
	fw.Log.WithFields(logrus.Fields{
		"account": account,
		"psw":  psw,
	}).Info("HandleRegister")
	accountkey := "account:account:" + account
	exist, _ := u.c.Do("EXIST", accountkey)
	if exist == false {
		u.c.Do("INCR", "account:count")
		id, _:= u.c.Do("GET", "account:count")
		u.c.Do("HSET", accountkey, "id", id)
		u.c.Do("HSET", accountkey, "psw", psw)
		code = com.E_Success
	}else{
		code = com.E_LoginAccountExist
		fw.Log.Warn("E_LoginAccountExist")
	}
	return
}

func(u *User)HandleAuth(game *Game, account, psw string)(code int, loginkey string, err error){
	fw.Log.WithFields(logrus.Fields{
		"name": account,
		"psw":  psw,
	}).Info("HandleLogin")
	accountkey := "account:account:" + account
	exist, _ := u.c.Do("EXIST", accountkey)
	if exist == false {
		code = com.E_LoginAccountNotExist
	}else{
		id, _:= u.c.Do("HGET", accountkey, "id")
		pswvalue, _ := u.c.Do("HGET", accountkey, "psw")
		if pswvalue == psw {
			loginkey = game.GenLoginKey(id)
			code = com.E_Success
		}else{
			code = com.E_LoginPasswordIncorrect
			fw.Log.Warn("E_LoginPasswordIncorrect")
		}
	}
	return
}
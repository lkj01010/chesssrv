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

func(u *User)Init()(err error){
	//setup connection
	u.c, err = redis.Dial("tcp", cfg.RedisAddr(),
		redis.DialReadTimeout(1 * time.Second), redis.DialWriteTimeout(1 * time.Second))
	if err != nil {
		fw.Log.Error("redis.Dial error")
		return
	}

	//select db
	_, err = u.c.Do("SELECT", cfg.RedisDBs["global"])
	if err != nil {
		fw.Log.Error("select err")
	}

	//fill keys
	b, _ := u.c.Do("EXISTS", "account:count")
	fw.Log.WithFields(logrus.Fields{
		"b": b,
	}).Info("")
	if b != 1 {
		b, err = u.c.Do("SET", "account:count", 3)
		fw.Log.WithFields(logrus.Fields{
			"b": b,
		}).Info("set ok?")
		if err != nil {
			fw.Log.Error(err.Error())
		}
	}
	return
}

func (u *User)Exit(){
	u.c.Close()
}

func(u *User)HandleRegister(account, psw string)(code int, err error){
	accountkey := "account:account:" + account
	fw.Log.WithField("accountkey", accountkey).Debug("")
	exist, _ := u.c.Do("EXISTS", accountkey)
	if exist == "0" {
		_, err = u.c.Do("INCR", "account:count")
		if err != nil{
			fw.Log.Error(err.Error())
		}
		id, _:= u.c.Do("GET", "account:count")
		u.c.Do("HSET", accountkey, "id", id)
		u.c.Do("HSET", accountkey, "psw", psw)
		code = com.E_Success
		fw.Log.Debug("Register success")
	}else{
		code = com.E_LoginAccountExist
		fw.Log.Warn("E_LoginAccountExist")
	}
	fw.Log.WithFields(logrus.Fields{
		"account": account,
		"psw":  psw,
		"exist": exist,
	}).Info("HandleRegister")
	return
}

func(u *User)HandleAuth(game *Game, account, psw string)(code int, loginkey string, err error){
	fw.Log.WithFields(logrus.Fields{
		"name": account,
		"psw":  psw,
	}).Info("HandleLogin")
	accountkey := "account:account:" + account
	exist, _ := u.c.Do("EXISTS", accountkey)
	if exist == 0 {
		code = com.E_LoginAccountNotExist
	}else{
		_id, _:= u.c.Do("HGET", accountkey, "id")
		id := _id.(string)
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
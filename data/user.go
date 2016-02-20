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

func (u *User)Init() (err error) {
	//setup connection
	u.c, err = redis.Dial("tcp", cfg.RedisAddr(),
		redis.DialReadTimeout(1 * time.Second), redis.DialWriteTimeout(1 * time.Second))
	if err != nil {
		fw.Log.Error("[user:Init] redis.Dial error")
		return
	}

	//select db
	s, err := u.c.Do("SELECT", cfg.RedisDBs["global"])
	if err != nil {
		return err
	}
	fw.Log.WithField("s", s).Error("select")
	fw.PrintType(s, "s")

	//fill keys
	b, _ := redis.Bool(u.c.Do("EXISTS", "account:count"))
	fw.Log.WithFields(logrus.Fields{
		"b": b,
	}).Info("fill key account")
	//	switch b.(type) {
	//	case interface{}:
	//		fw.Log.Debug("interface")
	//	case []byte:
	//		fw.Log.Debug("byte")
	//	case string:
	//		fw.Log.Debug("string")
	//	case *int:
	//		fw.Log.Debug("int")
	//	default:
	//		fw.Log.Debug("other")
	//	}

	if b == false {
		u.c.Do("SET", "account:count", 0)
	}
	return
}

func (u *User)Exit() {
	u.c.Close()
}
const (
	AccountAccountPre = "account:account:"
	AccountCountKey = "account:count"
	IdKey = "id"
	PswKey = "psw"
)
func (u *User)HandleRegister(account, psw string) (code int, err error) {
	accountkey := AccountAccountPre + account
	fw.Log.WithField("accountkey", accountkey).Debug("")
	exists, _ := redis.Bool(u.c.Do("EXISTS", accountkey))
	if !exists {
		u.c.Do("INCR", AccountCountKey)
		id, _ := u.c.Do("GET", AccountCountKey)
		u.c.Do("HSET", accountkey, IdKey, id)
		u.c.Do("HSET", accountkey, PswKey, psw)
		code = com.E_Success
		fw.Log.Debug("Register success")
	}else {
		code = com.E_LoginAccountExist
		fw.Log.Debug("E_LoginAccountExist")
	}
	fw.Log.WithFields(logrus.Fields{
		"account": account,
		"psw":  psw,
		"exist": exists,
	}).Info("[user:HandleRegister]")
	return
}

func (u *User)HandleAuth(game *Game, account, psw string) (code int, loginkey string, err error) {
	fw.Log.WithFields(logrus.Fields{
		"account": account,
		"psw":  psw,
	}).Debug("[user:HandleAuth]")
	accountkey := AccountAccountPre + account
	exists, _ := redis.Bool(u.c.Do("EXISTS", accountkey))
	if exists == false {
		code = com.E_LoginAccountNotExist
		fw.Log.Debug("[user:HandleAuth]E_LoginAccountNotExist")
	}else {
		id, _ := redis.String(u.c.Do("HGET", accountkey, IdKey))
		//		fw.PrintType(id, "id")
		pswvalue, _ := redis.String(u.c.Do("HGET", accountkey, PswKey))
		if pswvalue == psw {
			loginkey = game.GenLoginKey(id)
			code = com.E_Success
		}else {
			code = com.E_LoginPasswordIncorrect
			fw.Log.Warn("E_LoginPasswordIncorrect")
		}
	}
	return
}
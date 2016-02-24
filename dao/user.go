package dao
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

var user *User

func init() (err error) {
	//setup connection
	user = new(User)
	user.c, err = redis.Dial("tcp", cfg.RedisAddr(),
		redis.DialReadTimeout(1 * time.Second), redis.DialWriteTimeout(1 * time.Second))
	if err != nil {
		fw.Log.Error("[user:Init] redis.Dial error")
		return
	}

	//select db
	s, err := user.c.Do("SELECT", cfg.RedisDBs[cfg.Pf])
	if err != nil {
		return err
	}
	fw.Log.WithField("s", s).Error("select")
	fw.PrintType(s, "s")

	//fill keys
	b, _ := redis.Bool(user.c.Do("EXISTS", "account:count"))
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
		user.c.Do("SET", "account:count", 0)
	}
	return
}

func (u *User)exit() {
	u.c.Close()
}

const (
	AccountAccountPre = "account:account:"
	AccountCountKey = "account:count"
	IdKey = "id"
	PswKey = "psw"
)

type RegisterArgs struct {
	Account, Psw string
}

func (u *User)HandleRegister(args *RegisterArgs, code *int) error {
	accountkey := AccountAccountPre + args.Account
	fw.Log.WithField("accountkey", accountkey).Debug("")
	exists, _ := redis.Bool(u.c.Do("EXISTS", accountkey))
	if !exists {
		u.c.Do("INCR", AccountCountKey)
		id, _ := u.c.Do("GET", AccountCountKey)
		u.c.Do("HSET", accountkey, IdKey, id)
		u.c.Do("HSET", accountkey, PswKey, args.Psw)
		code = com.E_Success
		fw.Log.Debug("Register success")
	}else {
		code = com.E_LoginAccountExist
		fw.Log.Debug("E_LoginAccountExist")
	}
	fw.Log.WithFields(logrus.Fields{
		"account": args.Account,
		"psw":  args.Psw,
		"exist": exists,
	}).Info("[user:HandleRegister]")
	return nil
}

type AuthArgs struct {
	Account, Psw string
}

type AuthReply struct {
	Code int
	LoginKey string
}

func (u *User)HandleAuth(args *AuthArgs, reply *AuthReply) (err error) {
	fw.Log.WithFields(logrus.Fields{
		"account": args.Account,
		"psw":  args.Psw,
	}).Debug("[user:HandleAuth]")
	accountkey := AccountAccountPre + args.Account
	exists, _ := redis.Bool(u.c.Do("EXISTS", accountkey))
	if exists == false {
		reply.Code = com.E_LoginAccountNotExist
		fw.Log.Debug("[user:HandleAuth]E_LoginAccountNotExist")
	}else {
		id, _ := redis.String(u.c.Do("HGET", accountkey, IdKey))
		//		fw.PrintType(id, "id")
		pswvalue, _ := redis.String(u.c.Do("HGET", accountkey, PswKey))
		if pswvalue == args.Psw {
			reply.LoginKey = game.genLoginKey(id)
			reply.Code = com.E_Success
		}else {
			reply.Code = com.E_LoginPasswordIncorrect
			fw.Log.Warn("E_LoginPasswordIncorrect")
		}
	}
	return
}
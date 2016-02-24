package dao
import (
	"github.com/garyburd/redigo/redis"
	"chess/fw"
	"chess/com"
	"chess/cfg"
	"time"
	"github.com/lkj01010/log"
)

const (
	AccountAccountPre = "account:account:"
	AccountCountKey = "account:count"
	IdKey = "id"
	PswKey = "psw"
)

type User struct {
	c redis.Conn
}

var UserInst *User

func init() {
	//setup connection
	UserInst = new(User)
	var err error
	UserInst.c, err = redis.Dial("tcp", cfg.RedisAddr(),
		redis.DialReadTimeout(1 * time.Second), redis.DialWriteTimeout(1 * time.Second))
	if err != nil {
		log.Error("[user:Init] redis.Dial error")
		return
	}

	//select db
	s, err := UserInst.c.Do("SELECT", cfg.RedisDBs[cfg.Pf])
	if err != nil {
		return
	}
	log.Error("select")
	fw.PrintType(s, "s")

	//fill keys
	b, _ := redis.Bool(UserInst.c.Do("EXISTS", AccountCountKey))
	log.Info("fill key account")
	//	switch b.(type) {
	//	case interface{}:
	//		log..Debug("interface")
	//	case []byte:
	//		log..Debug("byte")
	//	case string:
	//		log..Debug("string")
	//	case *int:
	//		log..Debug("int")
	//	default:
	//		log..Debug("other")
	//	}

	if b == false {
		UserInst.c.Do("SET", AccountCountKey, 0)
	}
	return
}

func (u *User)exit() {
	u.c.Close()
}

type RegisterArgs struct {
	Account, Psw string
}

func (u *User)HandleRegister(args *RegisterArgs, reply *fw.RpcReply) error {
	accountkey := AccountAccountPre + args.Account
	exists, _ := redis.Bool(u.c.Do("EXISTS", accountkey))
	if !exists {
		u.c.Do("INCR", AccountCountKey)
		id, _ := u.c.Do("GET", AccountCountKey)
		u.c.Do("HSET", accountkey, IdKey, id)
		u.c.Do("HSET", accountkey, PswKey, args.Psw)
		reply.Code = com.E_Success
		log.Debug("Register success")
	}else {
		reply.Code = com.E_LoginAccountExist
		log.Debug("E_LoginAccountExist")
	}

	return nil
}

type AuthArgs struct {
	Account, Psw string
}

type AuthReply struct {
	Code     int
	LoginKey string
}

func (u *User)HandleAuth(args *AuthArgs, reply *AuthReply) (err error) {
	accountkey := AccountAccountPre + args.Account
	exists, _ := redis.Bool(u.c.Do("EXISTS", accountkey))
	if exists == false {
		reply.Code = com.E_LoginAccountNotExist
		log.Debug("[user:HandleAuth]E_LoginAccountNotExist")
	}else {
		id, _ := redis.String(u.c.Do("HGET", accountkey, IdKey))
		//		fw.PrintType(id, "id")
		pswvalue, _ := redis.String(u.c.Do("HGET", accountkey, PswKey))
		if pswvalue == args.Psw {
			reply.LoginKey = GameInst.genLoginKey(id)
			reply.Code = com.E_Success
		}else {
			reply.Code = com.E_LoginPasswordIncorrect
			log.Warning("E_LoginPasswordIncorrect")
		}
	}
	return
}
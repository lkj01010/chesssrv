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
	k_account_account_ = "account:account:"
	k_account_count = "account:count"
	k_id = "id"
	k_psw = "psw"
)

type User struct {
	model
}

func NewUser(p *Models) *User {
	u := new(User)
	c, err := redis.Dial("tcp", cfg.RedisAddr(),
		redis.DialReadTimeout(1 * time.Second), redis.DialWriteTimeout(1 * time.Second))
	if err != nil {
		log.Error(err.Error())
		return nil
	}
	u.c = c

	//select db
	s, err := c.Do("SELECT", cfg.RedisDBs[cfg.Pf])
	if err != nil {
		log.Error(err.Error())
		return nil
	}
	//temp
	fw.PrintType(s, "s")

	//fill keys
	b, _ := redis.Bool(c.Do("EXISTS", k_account_count))

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
		c.Do("SET", k_account_count, 0)
		log.Info("fill key:", k_account_count)
	}

	//register model
	u.parent = p
	return u
}

func (u *User)exit() {
	u.c.Close()
}

type UserRegisterArgs struct {
	Account, Psw string
}

func (u *User)HandleRegister(args *UserRegisterArgs, reply *fw.RpcReply) error {
	accountkey := k_account_account_ + args.Account
	exists, _ := redis.Bool(u.c.Do("EXISTS", accountkey))
	if !exists {
		u.c.Do("INCR", k_account_count)
		id, _ := u.c.Do("GET", k_account_count)
		u.c.Do("HSET", accountkey, k_id, id)
		u.c.Do("HSET", accountkey, k_psw, args.Psw)
		reply.Code = com.E_Success
		log.Debug("Register success")
	}else {
		reply.Code = com.E_LoginAccountExist
		log.Debug("E_LoginAccountExist")
	}

	return nil
}

type UserAuthArgs struct {
	Account, Psw string
}

type UserAuthReply struct {
	Code     int
	LoginKey string
}

func (u *User)HandleAuth(args *UserAuthArgs, reply *UserAuthReply) (err error) {
	accountkey := k_account_account_ + args.Account
	exists, _ := redis.Bool(u.c.Do("EXISTS", accountkey))
	if exists == false {
		reply.Code = com.E_LoginAccountNotExist
		log.Info("E_LoginAccountNotExist")
	}else {
		id, _ := redis.String(u.c.Do("HGET", accountkey, k_id))
		//		fw.PrintType(id, "id")
		pswvalue, _ := redis.String(u.c.Do("HGET", accountkey, k_psw))
		if pswvalue == args.Psw {
			reply.LoginKey = u.parent.Game.genLoginKey(id)
			reply.Code = com.E_Success
		}else {
			reply.Code = com.E_LoginPasswordIncorrect
			log.Info("E_LoginPasswordIncorrect")
		}
	}
	return
}
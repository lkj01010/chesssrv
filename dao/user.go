package dao
import (
	"github.com/garyburd/redigo/redis"
	"chess/fw"
	"chess/com"
	"chess/cfg"
	"time"
	"github.com/lkj01010/log"
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

type User_RegisterArgs struct {
	Account, Psw string
}

func (u *User)Register(args *User_RegisterArgs, reply *Reply) error {
	if args.Psw == "" {
		reply.Code = com.E_AgentPasswordCannotBeNull
	} else {
		accountkey := k_account_user_ + args.Account
		account, _ := redis.String(u.c.Do("GET", accountkey))
		//	exists, _ := redis.Bool(u.c.Do("EXISTS", accountkey))
		//	if !exists {
		if account == "" {
			// create id
			u.c.Do("INCR", k_account_count)
			id, _ := redis.String(u.c.Do("GET", k_account_count))

			// save account
			u.c.Do("SET", accountkey, id)
			u.c.Do("SADD", k_account_userlist, id)

			// save user
			userkey := k_user_ + string(id)
			u.c.Do("HSET", userkey, k_psw, args.Psw)
			u.c.Do("HSET", userkey, k_coin, 0)
			u.c.Do("HSET", userkey, k_nickname, "nickname")

			reply.Code = com.E_Success
			log.Debug("Register success")
		} else {
			reply.Code = com.E_AgentAccountExist
			log.Debug("E_AgentAccountExist")
		}
	}
	return nil
}

type User_AuthArgs struct {
	Account, Psw string
}

type User_AuthReply struct {
	Code int
	Id   string
}

func (u *User)Auth(args *User_AuthArgs, reply *User_AuthReply) (err error) {
	accountkey := k_account_user_ + args.Account
	//	exists, _ := redis.Bool(u.c.Do("EXISTS", accountkey))

	var id string
	log.Debugf("Auth:get %+v", accountkey)
	id, _ = redis.String(u.c.Do("GET", accountkey))
	// id may be ""
	psw, _ := redis.String(u.c.Do("HGET", k_user_ + id, k_psw))
	//	if id == "" {
	//		reply.Code = com.E_AgentAccountNotExist
	//	}else{
	//		reply.Code = com.E_Success
	//	}

	log.Debugf("Auth:id=%+v psw=%+v", id, psw)
	if id == "" || psw == "" {
		reply.Code = com.E_AgentAccountNotExist
	}else {
		if psw == args.Psw {
			reply.Code = com.E_Success
			reply.Id = id
		}else {
			reply.Code = com.E_AgentPasswordIncorrect
		}
	}

	//	if exists == false {
	//		reply.Code = com.E_AgentAccountNotExist
	//		log.Info("E_LoginAccountNotExist")
	//	}else {
	//		id, _ := redis.String(u.c.Do("HGET", accountkey, k_id))
	//		//		fw.PrintType(id, "id")
	//		pswvalue, _ := redis.String(u.c.Do("HGET", accountkey, k_psw))
	//		if pswvalue == args.Psw {
	//			reply.LoginKey = u.parent.Game.genLoginKey(id)
	//			reply.Code = com.E_Success
	//		}else {
	//			reply.Code = com.E_AgentPasswordIncorrect
	//			log.Info("E_AgentPasswordIncorrect")
	//		}
	//	}
	return
}
///////////////////////////////////////////////////////
type User_Info struct {
	Id       string    `json:"id"`
	Nickname string `json:"nickname"`
	Coin     int    `json:"coin"`
}

type User_InfoReply struct {
	Code int
	Info User_Info
}

func (u *User)Info(args *Args, reply *User_InfoReply) (err error) {
	var nickname string

	coin, _ := redis.Int(u.c.Do("HGET", k_user_ + args.Id, k_coin))
	nickname, _= redis.String(u.c.Do("HGET", k_user_ + args.Id, k_nickname))
	log.Infof("coin=%+v", coin)
	// fixme:这里用是否为空判断用户存在与否
	if nickname == "" {
		reply.Code = com.E_AgentAccountNotExist
	} else {
		reply.Code = com.E_Success
		reply.Info.Id = args.Id
		reply.Info.Nickname = nickname
		reply.Info.Coin = coin
	}
	return
}

///////////////////////////////////////////////////////
func (u *User)GetCoin(args *Args, reply *Reply) (err error) {
	var coin int
	coin, err = redis.Int(u.c.Do("HGET", k_user_ + args.Id, k_coin))
	if err != nil {
		log.Warningf("Coin:not found, id=%+v", args.Id)
		reply.Code = com.E_ValueNotFound
	} else {
		reply.Code = com.E_Success
		reply.Int = coin
	}
	return
}

///////////////////////////////////////////////////////
func (u *User)AddCoin(args *Args, reply *Reply) (err error) {
	_, err = u.c.Do("HINCRBY", k_user_ + args.Id, k_coin, args.Int)
	if err != nil {
		log.Warning("AddCoin:err=", err.Error())
		reply.Code = com.E_ValueNotFound
	} else {
		reply.Code = com.E_Success
	}
	return
}

package game
import (
	"encoding/json"
	"chess/com"
	"github.com/lkj01010/log"
	"chess/dao"
)

type playerAgent struct {
	connId   int
	id       string
	c        chan string
	sendFunc func(int, string)

	toGame chan string		// 发往game的channel

	info com.UserInfo
}

func NewPlayerAgent(connId int, sf func(int, string)) *playerAgent {
	return &playerAgent{
		connId: connId,
		c: make(chan string, 10),
		sendFunc: sf,
		toGame: make(chan string, 10),
	}
}

func (p *playerAgent)Go() {
	// get user info
	var reply dao.User_InfoReply
	modelInst.dao.Call("User.GetInfo", &dao.Args{Id: p.id}, &reply)
	p.info = reply.Info

	for {
		select {
		case rcv := <-p.c:
			p.handle(rcv)
		}
	}
}

func (p *playerAgent)Receive(msg string) {
	p.c <- msg
}

func (p *playerAgent)Send(msg string) {
	p.sendFunc(p.connId, msg)
}

func (p *playerAgent)handle(msgstr string) (err error) {
	var msg com.Msg
	if err = json.Unmarshal([]byte(msgstr), &msg); err != nil {
		return
	}

	switch com.Command(msg.Cmd) {
	case com.Cmd_Game_EnterReq:
		err = p.handleEnterReq(msg.Content)
	default:
		p.toGame <-msg.Content
	}
	return
}

func (p *playerAgent)handleEnterReq(contentstr string) (err error) {
	_, ok := modelInst.playerGames[p.connId]
	if ok {
		// 已经在游戏中，报错
		err = com.ErrAlreadyInGame
		log.Warning("game:model:handle:player enter req, err=", err.Error())
		return
	} else {
		var content EnterGameReq
		if err = json.Unmarshal([]byte(contentstr), &content); err != nil {
			return
		}
		p.id = content.Id

		//判断钱是否够
		isCoinEnough := (p.info.Coin >= RoomEnterCoin[content.RoomType])

		// 够入场费
		if isCoinEnough {
			// 塞进房间
			rt := RoomType(content.RoomType)
			if rt.IsValid() {
				game := modelInst.GetFreeGameByType(rt)
				game.PlayerEnter(content.Id, p.info, p.toGame, p.Send)
			} else {
				log.Warningf("game:model:handle:roomtype invalid, rt=%+v", content.RoomType)
				err = com.ErrRoomTypeInvalid
				return
			}
		}else {
			// 不够入场费
			// 返回response
			resp := com.MakeMsgString(com.Cmd_Game_EnterResp, com.E_CoinNotEnough, nil)
			p.Send(resp)
			return
		}
	}
	return
}
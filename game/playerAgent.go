package game
import (
	"chess/game/cow"
	"encoding/json"
"chess/com"
"github.com/lkj01010/log"
)

type playerAgent struct {
	connId int
	c chan string
	sendFunc func(int, string)
	game *cow.Game
}

func NewPlayerAgent(connId int, sf func(string)) *playerAgent {
	return &playerAgent{
		connId: connId,
		c: make(chan string, 10),
		sendFunc: sf,
	}
}

func (p *playerAgent)Go(){
	for {
		select {
		case rcv := <- p.c:
			p.handle(rcv)
		}
	}
}

func (p *playerAgent)Recive(msg string){
	p.c <- msg
}

func (p *playerAgent)Send(msg string) {
	p.sendFunc(msg)
}

func (p *playerAgent)handle(msg string) (err error) {
	var msg com.Msg
	if err = json.Unmarshal([]byte(msg.Content), &msg); err != nil {
		return
	}

	switch msg.Cmd {
	case Cmd_Game_EnterReq:
		_, e := modelInst.playerGames[p.connId]
		if e == nil {
			// 已经在游戏中，报错
			err = com.ErrAlreadyInGame
			log.Warning("game:model:handle:player enter req, err=", err.Error())
			return
		} else {
			var content EnterGame
			if err = json.Unmarshal([]byte(msg.Content), &content); err != nil {
				return
			}

			//判断钱是否够
			modelInst.dao.

			games := modelInst.freeGames[RoomType(content.RoomType)]
			// 塞进第一个房间
			var game *cow.Game
			if len(games) == 0 {
				// create game
				game = cow.NewGame(RoomEnterCoin[RoomType(content.RoomType)])
				game.Go()

				game.PlayerEnter(content.Id)
			} else {
				// put to first game
				game = games[0]
			}
		}
	case cow.Cmd_Cow_Ready:
	}
	return
}
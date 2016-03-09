package fw
import (
	log "github.com/lkj01010/log"
	"time"
)

const (
	CtrlRemoveAgent = "removeAgent"
)

type Model interface {
	Enter()
	Handle(req string) (resp string, err error)
	Exit()
	Hook(*Agent)
}

type Agent struct {
	Model
	ReadWriteCloser
	ConnId int
	out    chan string
	send   chan string
}

func NewAgent(m Model, rwc ReadWriteCloser, connId int) *Agent {
	a := &Agent{
		Model: m,
		ReadWriteCloser: rwc,
		ConnId: connId,
		out: make(chan string, 10),
		send: make(chan string, 10),
	}
	m.Hook(a)
	return a
}

func (a *Agent)Serve() (err error) {
	a.Enter()
	defer a.Exit()

	session := make(chan string, 5)
	go func(c chan string) {
		var buf string
		for {
			err = a.Read(&buf)
			if err != nil {
				//fixme: 通过正常流程退出
				log.Debug("agent read: ", err.Error())
				return
			}
			c <- buf
		}
	}(session)

	timeout := time.NewTimer(0)
	L:    for {
		timeout.Reset(10 * time.Second)
		select {

		case cmd := <-a.out:
			if cmd == CtrlRemoveAgent {
				log.Debug("recv ctrl: CtrlRemoveAgent")
				break L
			}

		case msg := <-session:
			var resp string
			resp, err = a.Handle(msg)
			if err != nil {
				log.Error("agent session: ", err.Error())
				return
			}
			a.Write(resp)

		case msg := <-a.send:
			if err = a.Write(msg); err != nil {
				// 写错误
				log.Error("client write failed: ", err.Error())
				return
			}

		case <-timeout.C:
			log.Debug("recv timeout")
			break L

		}
	}

	//	var buf, resp string
	//	err = a.Read(&buf):
	//	if err != nil {
	//		log.Error("read error:", err.Error())
	//		return
	//	}
	//	resp, err = a.Handle(buf)
	//	if err != nil {
	//		log.Error("agent serve:", err.Error())
	//		return
	//	}
	//	a.Write(resp)

	return
}

func (a *Agent)Cmd(cmd string) {
	a.out <- cmd
}

func (a *Agent)Send(msg string) {
	a.send <- msg
}
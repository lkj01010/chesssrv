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
	Ctrl chan string
}

func NewAgent(h Model, rw ReadWriteCloser) *Agent {
	a := &Agent{h, rw, make(chan string, 2)}
	h.Hook(a)
	return a
	// todo: modify agent to 2 as above this

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
		case msg := <-session:
			log.Debug("recv msg")
			var resp string
			resp, err = a.Handle(msg)
			if err != nil {
				log.Error("agent session: ", err.Error())
				return
			}
			a.Write(resp)

		case <-timeout.C:
			log.Debug("recv timeout")
			break L
		case ctrl := <-a.Ctrl:
			if ctrl == CtrlRemoveAgent {
				log.Debug("recv ctrl: CtrlRemoveAgent")
				break L
			}
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
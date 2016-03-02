package fw
import (
	log "github.com/lkj01010/log"
	"time"
)

type ReadWriteCloser interface {
	Read(msg *string) error
	Write(msg string) error
	Close() error
}

type Handler interface {
	Handle(req string) (resp string, err error)
}

type Agent struct {
	Handler
	ReadWriteCloser
	ctrl chan string
}

func NewAgent(h Handler, rw ReadWriteCloser) *Agent {
	return &Agent{h, rw, make(chan string, 0)}
}

func (a *Agent)Serve() (err error) {
	session := make(chan string, 0)
	go func(c chan string) {
		var buf string
		for {
			err = a.Read(&buf)
			if err != nil {
				log.Debug("agent read: ", err.Error())
				return
			}
			c <-buf
		}
	}(session)

	timeout := time.NewTimer(0)
L:	for {
		timeout.Reset(5 * time.Second)
		select {
		case msg := <-session:
			var resp string
			resp, err = a.Handle(msg)
			if err != nil {
				log.Error("agent session: ", err.Error())
				return
			}
			a.Write(resp)
			log.Debug("recv msg")
		case <-timeout.C:
			log.Debug("recv timeout")
			break L
		case <-a.ctrl:
			log.Debug("recv ctrl")
			break;
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
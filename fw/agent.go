package fw
import (
	log "github.com/lkj01010/log"
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
}

func NewAgent(h Handler, rw ReadWriteCloser) *Agent {
	return &Agent{h, rw}
}

func (a *Agent)Serve() (err error) {
	var buf, resp string
	for {
		err = a.Read(&buf)
		if err != nil {
			log.Error("read error:", err.Error())
			return
		}
		//		fmt.Println("id=" + ra.Id() + ", read=" + buf)
		resp, err = a.Handle(buf)
		if err != nil {
			log.Error("handle error:", err.Error())
			return
		}
		//		if resp != "" {
		a.Write(resp)
		//		}else{
		//			a.Write("recive ''")
		//		}
	}
	return
}

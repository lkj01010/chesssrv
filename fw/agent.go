package fw
import (
	log "github.com/lkj01010/log"
)

type ReadWriter interface {
	Read(msg *string) (err error)
	Write(msg string) (err error)
}

type Handler interface {
	Handle(req string) (resp string, err error)
}

type Agent struct {
	Handler
	ReadWriter
}

func NewAgent(h Handler, rw ReadWriter) *Agent {
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

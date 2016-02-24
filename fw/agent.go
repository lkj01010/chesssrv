package fw
import (
	log "github.com/lkj01010/log"
)

type ReadWriter interface {
	Read(msg *string) (err error)
	Write(msg string) (err error)
}

type Agent interface {
	Handle(req string) (resp string, err error)
}

type IpcAgent struct {
	Agent
	ReadWriter
}

func NewIpcAgent(agent Agent, rw ReadWriter) *IpcAgent {
	return &IpcAgent{agent, rw}
}

func (a *IpcAgent)Serve() (err error) {
	for {
		var buf string
		err = a.Read(&buf)
		if err != nil {
			goto Error
		}
		//		fmt.Println("id=" + ra.Id() + ", read=" + buf)
		resp, err := a.Handle(buf)
		if err != nil {
			goto Error
		}
		//		if resp != "" {
		a.Write(resp)
		//		}else{
		//			a.Write("recive ''")
		//		}
	}
	return

	Error:
	log.Error("[agent:Server]" + err.Error())
	return
}

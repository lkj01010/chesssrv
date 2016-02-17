package login
import "chess/fw"

type ReadWriter interface {
	Read(msg *string) (err error)
	Write(msg string)(err error)
}

type Agent struct {
	rw ReadWriter
}

func NewAgent(rw ReadWriter) (a *Agent) {
	fw.Log.Info("login:NewAgent")
	return &Agent{rw}
}
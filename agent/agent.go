package agent
import (
	"fmt"
)

type ReadWriter interface {
	Read(msg *string) (err error)
	Write(msg string)(err error)
}

type Agent struct {
	rw ReadWriter
	id string
}

func NewAgent(rw ReadWriter, id string) (a *Agent){
	fmt.Println("NewAgent, id=" + id)
	return &Agent{rw, id}
}

func (a *Agent)Id()(string){
	return a.id
}

func (a *Agent)Handle(msg string)(err error){
	reply := a.id + `:` + msg

	err = a.Write(reply)
	return
}

func (a *Agent)Write(msg string)(error){
	return a.rw.Write(msg)
}

func (a *Agent)Read(msg *string)(error){
	return a.rw.Read(msg)
}
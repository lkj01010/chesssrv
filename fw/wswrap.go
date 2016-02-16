package fw
import (
	"fmt"
	"golang.org/x/net/websocket"
)

type WsReadWriter struct {
	ws *websocket.Conn
}

func (w *WsReadWriter)Read(msg *string) (err error) {
	err = websocket.Message.Receive(w.ws, msg)
	fmt.Printf("recv:%#v\n", *msg)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func (w *WsReadWriter)Write(msg string) (err error) {
	err = websocket.Message.Send(w.ws, msg)
	fmt.Printf("send:%#v\n", msg)
	if err != nil {
		fmt.Println(err)
	}
	return
}

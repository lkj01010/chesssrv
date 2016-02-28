package fw
import (
	"fmt"
	"golang.org/x/net/websocket"
)

type WsClient struct {
	ws *websocket.Conn
}

func NewWsReadWriter(ws *websocket.Conn) *WsClient {
	return &WsClient{ws}
}

func (w *WsClient)Read(msg *string) (err error) {
	err = websocket.Message.Receive(w.ws, msg)
	fmt.Printf("recv: %#v\n", *msg)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func (w *WsClient)Write(msg string) (err error) {
	err = websocket.Message.Send(w.ws, msg)
	fmt.Printf("send: %#v\n", msg)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func (w *WsClient)Close() (err error) {
	err = w.ws.Close()
	return
}

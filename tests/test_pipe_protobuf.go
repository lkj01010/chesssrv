package main
import (
	"net/http"
	"golang.org/x/net/websocket"
	log "github.com/lkj01010/log"
	"fmt"
	"github.com/golang/protobuf/proto"
"chess/com"
)

func makeProtoString() []byte {
	msg := &com.UserInfo{
		Id:  proto.String("lkj"),
		Nickname: proto.String("mid"),
	} //m
	buffer, err := proto.Marshal(msg)
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		return []byte("err")
	}
	fmt.Println("makeProtoString=", buffer)
	return buffer
}

func testSendRecvServer(ws *websocket.Conn) {
	fmt.Printf("sendRecvServer %#v\n", ws)
	for {
		var buf string
		// Receive receives a text message from client, since buf is string.
		err := websocket.Message.Receive(ws, &buf)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("recv:%q\n", buf)

////		buffer := make([]byte, len(buf))
//		buffer := make([]byte, 1000)
//		msg := &com.UserInfo{}
//		err = proto.Unmarshal(buffer, msg) //unSerialize
//		if err != nil {
//			fmt.Printf("failed: %s\n", err)
//			return
//		}
//		fmt.Printf("read: %+v", msg)

		// Send sends a text message to client, since buf is string.
		byteBuf := makeProtoString()
		err = websocket.Message.Send(ws, string(byteBuf))
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("send:%q\n", byteBuf)



		buffer := make([]byte, 1000)
		msg := &com.UserInfo{}
		err = proto.Unmarshal(buffer, msg) //unSerialize
		if err != nil {
			fmt.Printf("failed: %s\n", err)
			return
		}
		fmt.Printf("parse to: %+v", msg)

	}
	fmt.Println("sendRecvServer finished")
}

func testReadWriteServer(ws *websocket.Conn) {
	fmt.Printf("readWriteServer %#v\n", ws.Config())
	for {
		buf := make([]byte, 1000)
		// Read at most 100 bytes.  If client sends a message more than
		// 100 bytes, first Read just reads first 100 bytes.
		// Next Read will read next at most 100 bytes.
		n, err := ws.Read(buf)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("recv:%+v\n", buf[:n])
		// Write send a message to the client.
//		n, err = ws.Write(buf[:n])
		if err != nil {
			fmt.Println(err)
			break
		}
//		fmt.Printf("send:%+v\n", buf[:n])

		///////////
		byteBuf := makeProtoString()
		msg := &com.UserInfo{}
		err = proto.Unmarshal(byteBuf, msg) //unSerialize
		if err != nil {
			fmt.Printf("failed: %s\n", err)
			return
		}
		n, err = ws.Write(byteBuf)
		fmt.Printf("send:%+v len=%+v", byteBuf, n)
	}
	fmt.Println("readWriteServer finished")
}

func testSendRecvBinaryServer(ws *websocket.Conn) {
	fmt.Printf("sendRecvBinaryServer %#v\n", ws)
	for {
		var buf []byte
		// Receive receives a binary message from client, since buf is []byte.
		err := websocket.Message.Receive(ws, &buf)
		if err != nil {
			fmt.Println(err)
		}

		byteBuf := makeProtoString()
		msg := &com.UserInfo{}
		err = proto.Unmarshal(byteBuf, msg) //unSerialize
		if err != nil {
			fmt.Printf("failed: %s\n", err)
			return
		}

		// Send sends a binary message to client, since buf is []byte.
		err = websocket.Message.Send(ws, byteBuf)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("send:%#v\n", byteBuf)
	}
	fmt.Println("sendRecvBinaryServer finished")
}

func main() {
	port := "8080"
	http.Handle("/", websocket.Handler(testSendRecvBinaryServer))
	log.Info("server start on:", port)

	http.ListenAndServe(":" + port, nil)
}
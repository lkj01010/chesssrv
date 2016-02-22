package login
import (
	"testing"
	"golang.org/x/net/websocket"
	"fmt"
	"strconv"
	"chess/fw"
	"chess/data"
	"net/http"
	"net"
	"bytes"
	"sync"
	"net/http/httptest"
)

func serve(ws *websocket.Conn) {
	connCnt ++
	fmt.Printf("agent come, access cnt=%s\n", strconv.Itoa(connCnt))

	agent := fw.NewIpcAgent(&LoginServer{}, fw.NewWsReadWriter(ws))
	agent.Serve()
}
var (
	connCnt = 0
	dUser *data.User
)
func onInit() {
	dUser = new(data.User)
	dUser.Init()
}

func onExit() {
	dUser.Exit()
}

func startServer() {
	onInit()
	defer func() {
		onExit()
	}()

	http.Handle("/login", websocket.Handler(serve))
	//	http.ListenAndServe(":8000", nil)

	server := httptest.NewServer(nil)
	serverAddr = server.Listener.Addr().String()
	fw.Log.Print("Test WebSocket server listening on ", serverAddr)
}

var (
	serverAddr string
	once sync.Once
)

func newConfig(t *testing.T, path string) *websocket.Config {
	config, _ := websocket.NewConfig(fmt.Sprintf("ws://%s%s", serverAddr, path), "http://localhost")
	return config
}

func TestServer(t *testing.T) {
	once.Do(startServer)

	// websocket.Dial()
	client, err := net.Dial("tcp", serverAddr)
	if err != nil {
		t.Fatal("dialing", err)
	}
	conn, err := websocket.NewClient(newConfig(t, "/login"), client)
	if err != nil {
		t.Errorf("WebSocket handshake error: %v", err)
		return
	}

	for i := 0; i < 10; i++ {
		msg := []byte("hello, world")
		fw.PrintType(msg, "msg")
		msg = append(msg, byte(i))
		//		append(msg, []byte("\n"))
		if _, err := conn.Write(msg); err != nil {
			t.Errorf("Write: %v", err)
		}
		var actual_msg = make([]byte, 512)
		n, err := conn.Read(actual_msg)
		if err != nil {
			t.Errorf("Read: %v", err)
		}
		actual_msg = actual_msg[0:n]
		if !bytes.Equal(msg, actual_msg) {
			t.Logf("Test: send %q got %q", msg, actual_msg)
		}
	}

	conn.Close()
}

func TestServer2(t *testing.T){
	once.Do(startServer)

}
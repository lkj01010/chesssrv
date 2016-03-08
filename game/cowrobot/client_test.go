package cowrobot_test
import (
	"testing"
	"chess/cfg"
	"time"
	"chess/fw"
	"chess/game/cowrobot"
)

func TestClient1(t *testing.T){
	ctrl := make(chan string, 1)
	cli, err := fw.NewClient(cfg.AgentAddr(), &cowrobot.Clientmodel{}, ctrl)
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log(cli)
	go cli.Loop()


	s := string(`{"cmd":104,"content":"{\"account\":\"testUtf\",\"psw\":\"pswlk22\"}"}`)
	cli.Send(s)
	time.Sleep(10 * time.Second)
	t.Log("end")
}

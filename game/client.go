package game
import (
	"chess/cfg"
	"chess/fw"
)

type Client struct {
	*fw.RpcClient
}

func NewClient() (c *Client, err error) {
	var rc *fw.RpcClient
	rc, err = fw.NewRpcClient(cfg.GameAddr())
	c = &Client{rc}
	return
}

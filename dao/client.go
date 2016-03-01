package dao
import (
	"chess/cfg"
	"chess/fw"
)

//todo: remove this file
type Client struct {
	*fw.RpcClient
}

func NewClient() (c *Client, err error) {
	var rc *fw.RpcClient
	rc, err = fw.NewRpcClient(cfg.DaoAddr())
	c = &Client{rc}
	return
}


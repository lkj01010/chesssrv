package dao
import (
	"chess/cfg"
	"net/rpc"
)

//todo: del
type Client struct {
	rpc.Client
}

func NewClient() (c *Client, err error) {
	var cli *rpc.Client
	if cli, err = rpc.Dial("tcp", cfg.DaoAddr()); err != nil {
		return
	}
	c = &Client{*cli}
	return
}


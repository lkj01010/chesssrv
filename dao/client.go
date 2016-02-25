package dao
import (
"net/rpc"
"chess/cfg"
)

type Client struct {
	cli *rpc.Client
}

func NewClient() (cli *Client,err error){
	var cli *rpc.Client
	if cli, err = rpc.Dial("tcp",cfg.IPs[cfg.SrvType] + cfg.DaoPort); err != nil {
		return
	}
	return &Client{cli}, nil
}


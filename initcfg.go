package main

import (
	"chess/cfg"
	"github.com/lkj01010/log"
)

func main() {
	cfg.FlushCfgToDB()
	log.Info("FlushCfgToDB")
}

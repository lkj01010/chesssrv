package main

import (
	"chess/cfg"
	"chess/fw"
	"github.com/Sirupsen/logrus"
)

func main() {
	cfg.FlushCfgToDB()
	fw.Log.Info("Started observing beach")
	fw.Log.WithFields(logrus.Fields{
		"omg":    true,
		//		"err":    err,
		"number": 100,
	}).Warn("The ice breaks!")
}

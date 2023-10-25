package main

import (
	"github.com/dhurimkelmendi/pack_delivery_api/config"
	"github.com/dhurimkelmendi/pack_delivery_api/server"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Infof("Server starting ...")
	run()
}

func run() {
	config.GetDefaultInstance().SetLogLevel()
	config.GetDefaultInstance().LogConfigs()
	server.GetDefaultInstance().Start()
}

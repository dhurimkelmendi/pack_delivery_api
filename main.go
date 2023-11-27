package main

import (
	"os"

	"github.com/dhurimkelmendi/pack_delivery_api/config"
	"github.com/dhurimkelmendi/pack_delivery_api/db"
	"github.com/dhurimkelmendi/pack_delivery_api/migrations"
	"github.com/dhurimkelmendi/pack_delivery_api/server"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Infof("Server starting ...")

	if len(os.Args) > 1 {
		action := os.Args[1]

		if action == "migrate" {
			if len(os.Args) > 2 {
				migrate(os.Args[2])
			} else {
				logrus.Fatal("Missing migration action.")
			}
		} else {
			logrus.Fatalf("Unknown action: %s", action)
		}
	} else {
		run()
	}
	run()
}

func run() {
	config.GetDefaultInstance().SetLogLevel()
	config.GetDefaultInstance().LogConfigs()
	server.GetDefaultInstance().Start()
}
func migrate(action string) {
	logrus.Infof("Starting migration -- action: %s", action)
	config.GetDefaultInstance().SetLogLevel()
	config.GetDefaultInstance().LogConfigs()

	dbConn := db.GetDefaultInstance()

	if action == "reset" {
		migrations.Reset(dbConn.GetDB())
	} else {
		migrations.Migrate(action, dbConn.GetDB())
	}
}

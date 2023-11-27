package migrations

import (
	"github.com/go-pg/migrations/v8"
	"github.com/sirupsen/logrus"
)

func init() {
	migrations.Register(func(db migrations.DB) error {
		logrus.Infoln("Creating pack_sizes table")
		_, err := db.Exec(`
		CREATE TABLE pack_sizes (
			size int
		);`)
		return err
	}, func(db migrations.DB) error {
		logrus.Infoln("Dropping pack_sizes table")
		_, err := db.Exec(`
			DROP TABLE IF EXISTS pack_sizes CASCADE;
		`)
		return err
	})
}

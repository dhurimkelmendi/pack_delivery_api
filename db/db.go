package db

import (
	"crypto/tls"
	"fmt"

	"github.com/dhurimkelmendi/pack_delivery_api/config"
	"github.com/dhurimkelmendi/pack_delivery_api/models"

	"github.com/go-pg/pg/extra/pgdebug"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"

	//blank import pq
	_ "github.com/lib/pq"
)

// ErrNoMatch is returned when we request a row that doesn't exist
var ErrNoMatch = fmt.Errorf("no matching record")

// ErrUserForbidden is returned when the current user has no access to execute the command
var ErrUserForbidden = fmt.Errorf("user is forbidden")

// Database is a struct that contains references to db config and connection
type Database struct {
	config *config.Config
	db     *pg.DB
}

var defaultInstance *Database

// GetDefaultInstance returns the default instance of Database
func GetDefaultInstance() *Database {
	if defaultInstance == nil {
		defaultInstance = &Database{
			config: config.GetDefaultInstance(),
		}
		defaultInstance.connect()
		// register all many-to-many relationships
		orm.RegisterTable((*models.PackSize)(nil))
	}
	return defaultInstance
}

// GetDB returns the *pg.DB connection
func (d *Database) GetDB() *pg.DB {
	return d.db
}
func (d *Database) connect() {
	d.db = pg.Connect(&pg.Options{
		Addr:     d.config.DatabaseHost + ":" + d.config.DatabasePort,
		Database: d.config.DatabaseName,
		User:     d.config.DatabaseUsername,
		Password: d.config.DatabasePassword,
		TLSConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	})

	if d.config.DebugDatabase {
		// Print all queries.
		d.db.AddQueryHook(pgdebug.DebugHook{
			Verbose: true,
		})
	}
}

package database

import (
	"database/sql"
	"merlindex/example/config"
	"time"

	"github.com/tursodatabase/go-libsql"
	sqlite "github.com/ytsruh/gorm-libsql"
	"gorm.io/gorm"
)

func Db() (*gorm.DB, error) {
	connector, err := libsql.NewEmbeddedReplicaConnector(
		config.LocalDatabaseName,
		config.DatabaseUrl,
		libsql.WithAuthToken(config.DatabaseAuthToken),
		libsql.WithSyncInterval(time.Minute),
	)
	if err != nil {
		return nil, err
	}

	conn := sql.OpenDB(connector)
	return gorm.Open(sqlite.New(sqlite.Config{
		Conn: conn,
	}), &gorm.Config{})
}

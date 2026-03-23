package database

import (
	"database/sql"
	"merlindex/example/config"

	_ "github.com/mattn/go-sqlite3"
)

func Db() (*sql.DB, error) {
	return sql.Open("sqlite3", config.LocalDatabaseName)

	// connector, err := libsql.NewEmbeddedReplicaConnector(
	// 	config.LocalDatabaseName,
	// 	config.DatabaseUrl,
	// 	libsql.WithAuthToken(config.DatabaseAuthToken),
	// 	libsql.WithSyncInterval(time.Minute),
	// )
	// if err != nil {
	// 	return nil, err
	// }

	// return sql.OpenDB(connector), nil
}

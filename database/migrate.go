package database

import "os"

func Migrate() error {
	conn, err := Db()
	if err != nil {
		return err
	}

	sql, err := os.ReadFile("database/up.sql")
	if err != nil {
		return err
	}

	_, err = conn.Exec(string(sql))
	return err
}

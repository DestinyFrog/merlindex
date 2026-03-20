package database

func Migrate() error {
	conn, err := Db()
	if err != nil {
		return err
	}

	return conn.AutoMigrate()
}

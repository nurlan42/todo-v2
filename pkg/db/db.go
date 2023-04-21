package db

import (
	"fmt"
	"github.com/nurlan42/todo/config"

	"database/sql"
	_ "github.com/lib/pq"
)

func Connect(db config.Psql) (*sql.DB, error) {
	dbInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%v",
		db.Host, db.Port, db.UserName, db.Password, db.Name, db.SSLMode)

	sqlDB, err := sql.Open(db.Type, dbInfo)

	//defer sqlDB.Close()

	if err != nil {
		return nil, err
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}

	return sqlDB, nil
}

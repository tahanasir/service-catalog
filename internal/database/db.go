package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/tahanasir/service-catalog/internal/config"
)

func NewDBConn(c *config.Config) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", c.DatabaseUser, c.DatabasePassword, c.DatabaseHost, c.DatabasePort, c.DatabaseName)
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("db conn: unable to connect: %w", err)
	}

	db.SetMaxOpenConns(c.MaxOpenConns)
	db.SetMaxIdleConns(c.MaxIdleConns)

	return db, nil
}

package infrastructures

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

func NewSQLite(dbPath string) (*sql.DB, error) {
	// note: the busy_timeout pragma must be first because
	// the connection needs to be set to block on busy before WAL mode
	// is set in case it hasn't been already set by another connection
	pragmas := "_busy_timeout=10000&_journal_mode=WAL&_foreign_keys=1&_synchronous=NORMAL"

	conn, err := sql.Open("sqlite3", fmt.Sprintf("%s?%s", dbPath, pragmas))
	if err != nil {
		return nil, err
	}

	conn.SetMaxOpenConns(1000)
	conn.SetMaxIdleConns(30)
	conn.SetConnMaxIdleTime(5 * time.Minute)

	return conn, nil
}

func NewSQLiteWithGorm(dbPath string) (*gorm.DB, error) {

	gormDB, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return gormDB, nil

}

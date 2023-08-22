package wallet

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/file"
	"log"
	"wallet/application/infrastructures"
)

// TODO (ikhsan) : ganti migration nya ke folder documents
func (g *Wallet) Migration() *migrate.Migrate {
	sqlite, err := infrastructures.NewSQLite(fmt.Sprintf("%s/%s", g.baseDir, "wallet.db"))
	instance, err := sqlite3.WithInstance(sqlite, &sqlite3.Config{})
	if err != nil {
		log.Fatal(err)
	}

	fSrc, err := (&file.File{}).Open("./migrations")
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithInstance("file", fSrc, "sqlite3", instance)
	if err != nil {
		log.Fatal(err)
	}

	return m

}

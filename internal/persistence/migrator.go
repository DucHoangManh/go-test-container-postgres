package persistence

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	_ "github.com/lib/pq"
)

func MigrationUp(completeDsn string) error {
	iofsDriver, err := iofs.New(EmbeddedFiles, "migrations")
	if err != nil {
		return err
	}

	migrator, err := migrate.NewWithSourceInstance("iofs", iofsDriver, completeDsn)
	if err != nil {
		return err
	}

	return migrator.Up()
}

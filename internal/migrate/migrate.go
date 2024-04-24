package migrate

import (
	"errors"
	"fmt"
	"github.com/akram620/alif/internal/config"
	"github.com/golang-migrate/migrate/v4"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func ApplyMigrations(path string) error {
	m, err := migrate.New(fmt.Sprintf("file://%s", path), config.Values.DatabaseURL)
	if err != nil {
		return err
	}
	defer m.Close()

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}

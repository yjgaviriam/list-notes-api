package app

/*
import (
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"log"
)

func (a *App) Migrate() {
	driver, err := postgres.WithInstance(a.DB, &postgres.Config{})
	if err != nil {
		log.Println(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://./database/migrations/",
		"list-notes-db", driver)
	if err != nil {
		log.Println(err)
	}
	if err := m.Steps(2); err != nil {
		log.Println(err)
	}
}
*/

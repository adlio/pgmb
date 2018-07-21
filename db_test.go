package pgmb

import (
	"log"
	"os"

	"database/sql"
)

var DB *sql.DB

func init() {
	var err error

	dsn := os.Getenv("PGMB_TEST_DSN")
	if dsn == "" {
		log.Fatal("The PGMB_TEST_DSN variable is required to run tests.\nExample: PGMB_TEST_DSN=postgres://user:pass@localhost/db?sslmode=disable&search_path=musicbrainz,public")
	}
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
}

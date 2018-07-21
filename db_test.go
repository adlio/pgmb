package pgmb

import (
	"log"
	"os"

	"database/sql"
)

var TESTDB DB

func init() {
	var err error
	var db *sql.DB
	dsn := os.Getenv("PGMB_TEST_DSN")
	if dsn == "" {
		log.Fatal("The PGMB_TEST_DSN variable is required to run tests.\nExample: PGMB_TEST_DSN=postgres://user:pass@localhost/db?sslmode=disable&search_path=musicbrainz,public")
	}
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	TESTDB = NewDB(db)
}

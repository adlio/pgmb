# MusicBrainz PostgreSQL Package for Go

This package includes a set of functions to access data in the raw MusicBrainz
PostgreSQL database tables. This is a read-only package designed to make certain
data retrieval features easier than with the XML Web Service API the MusicBrainz
server provides.

## Usage Instructions

All functions accept a `pgmb.DB` as their first argument. This is a wrapper around the standard
library DB object. Use `NewDB(*sql.DB)` to create one:

```Go
db, err = sql.Open("postgres", "postgres://user:pass@localhost/db_name?sslmode=disable&search_path=musicbrainz,public")
if err != nil {
    log.Fatal(err)
}
mbDB := pgmb.NewDB(db)
```

## Test Instructions

Tests require a completely populated MusicBrainz database connection information must be supplied
as an environment variable `PGMB_TEST_DSN`.

Here's an example:

```Go
PGMB_TEST_DSN=postgres://user:pass@localhost/db_name?sslmode=disable&search_path=musicbrainz,public
```

## TODO

1. Test features which would allow looking up the "canonical" (earliest Official, U.S.-released occurrence)
of a recording based on an artist name and song title. This likely is a search feature on Track.
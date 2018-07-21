# MusicBrainz PostgreSQL Package for Go

This package includes a set of functions to access data in the raw MusicBrainz
PostgreSQL database tables. This is a read-only package designed to make certain
data retrieval features easier than with the XML Web Service API the MusicBrainz
server provides.

## Usage Instructions

All functions accept a `*sql.DB` as their first argument.

## Test Instructions

Tests require a completely populated MusicBrainz database connection information must be supplied
as an environment variable `PGMB_TEST_DSN`.

Here's an example:

```
PGMB_TEST_DSN=postgres://user:pass@localhost/database_name?sslmode=disable&search_path=musicbrainz,public
```
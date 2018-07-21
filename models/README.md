# MusicBrainz PostgreSQL Package for Go

This package includes a set of functions to access data in the raw MusicBrainz
PostgreSQL database tables. This is a read-only package designed to make certain
data retrieval features easier than with the XML Web Service API the MusicBrainz
server provides.

## Usage Instructions

All functions accept a `*sql.DB` as their first argument.
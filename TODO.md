# TODO

* Document the need for setting up GIN indexes:

    ```
    SET search_path=musicbrainz,public;

    CREATE INDEX IF NOT EXISTS idx_artist_lower_name_gin ON artist USING GIN (lower(name) gin_trgm_ops);

    CREATE INDEX IF NOT EXISTS idx_artist_alias_lower_name_gin ON artist_alias USING GIN (lower(name) gin_trgm_ops);

    CREATE INDEX IF NOT EXISTS idx_release_group_lower_name_gin ON release_group USING GIN (lower(name) gin_trgm_ops);
    ```
* Functions to search release groups
* Functions to search tracks / canonical tracks
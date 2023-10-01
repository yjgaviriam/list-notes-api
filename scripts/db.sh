#!/bin/bash

set -e
export PG_PASSWORD=postgres123;
psql -v ON_ERROR_STOP=1 --username "postgres" --dbname "list-notes-db" <<-EOSQL
  CREATE DATABASE list-notes-db;
  GRANT ALL PRIVILEGES ON DATABASE list-notes-db TO "postgres";
EOSQL
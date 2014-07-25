package app

import (
  // 3rd-party library
  _ "github.com/lib/pq"
  // Built-in
  "database/sql"
  // Badge
)

type Database struct {
  db *sql.DB
}

var connection *Database

// NOTE: For heroku, just use DSN to connect database
// NOTE: Implement singleton pattern to prevent create multi connection to database
func OpenConnection(dsn string) *Database {
  if connection == nil {
    db, err := sql.Open("postgres", dsn)
    if err != nil {
      panic(err) // If database connection failed, direct panic it
    }
    connection = &Database{ db: db }
  }
  return connection
}

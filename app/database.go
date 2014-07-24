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

// NOTE: For heroku, just use DSN to connect database
func OpenConnection(dsn string) *Database {
  db, err := sql.Open("postgres", dsn)
  if err != nil {
    panic(err) // If database connection failed, direct panic it
  }
  return &Database{ db: db }
}

type User struct {
  ID int
  Email string
}

func (conn *Database) GetUsers() []User {
  rows, err := conn.db.Query("SELECT id, email FROM users")
  // TODO: design error handler
  if err != nil {
    panic(err)
  }
  users := []User{}
  defer rows.Close()
  for rows.Next() {
    var id int
    var email string
    rows.Scan(&id, &email)
    users = append(users, User{ID: id, Email: email})
  }

  return users
}

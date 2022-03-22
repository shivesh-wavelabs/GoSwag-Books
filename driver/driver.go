package driver

import (
  "database/sql"
  "log"
  "os"

  "github.com/lib/pq"
)

var db *sql.DB

// Custom error function
func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Connection to the Postgres DB
func ConnectDB() *sql.DB {
  pgUrl, err := pq.ParseURL(os.Getenv("ELEPHANTSQL_URL"))
  logFatal(err)

  db, err = sql.Open("postgres", pgUrl)
  logFatal(err)

  err = db.Ping()
  logFatal(err)

  return db
}

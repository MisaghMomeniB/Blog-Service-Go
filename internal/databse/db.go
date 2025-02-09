package databse

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // PostgreSQL driver
)

// Connect connects to the database and returns the connection
func Connect() *sql.DB {
	connStr := "user=username dbname=mydb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

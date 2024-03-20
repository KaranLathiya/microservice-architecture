package database

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

var once sync.Once
var db *sql.DB

func Connect() (*sql.DB, error) {

	var err error
	once.Do(func() {
		// Connection parameters
		connStr := "postgresql://root@localhost:26257/expert?sslmode=disable"
		db, err = sql.Open("postgres", connStr)
		if err != nil {
			fmt.Println("Database Connection err", err)
			return
		}
		err = db.Ping()
		if err != nil {
			log.Fatal("Error pinging the database:", err)
		}
		fmt.Println("Successfully connected to CockroachDB!")
	})
	return db, err
}

func GetDB() *sql.DB {
	return db
}

package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	connStr := "host=db user=user password=password dbname=social sslmode=disable"

	var err error
	for i := 1; i <= 10; i++ {
		DB, err = sql.Open("postgres", connStr)
		if err != nil {
			log.Printf("Attempt %d: failed to open connection: %s", i, err)
		} else {
			err = DB.Ping()
			if err == nil {
				fmt.Println("✅ Connected to PostgreSQL")
				return
			}
			log.Printf("Attempt %d: ping failed: %s", i, err)
		}
		time.Sleep(2 * time.Second)
	}

	log.Fatal("❌ Could not connect to database after 10 attempts:", err)
}

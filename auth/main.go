package main

import (
	"auth/data"
	"database/sql"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

const webPort = "80"

var counts int64

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {
	log.Println("Starting authentication service")

	conn := connectToDB()
	app := Config{
		DB:     conn,
		Models: data.New(conn),
	}
	http.ListenAndServe(":8000", app.routes())
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func connectToDB() *sql.DB {
	// dsn := os.Getenv("DSN")
	dsn := "host=localhost port=5432 user=postgres dbname=users password=rafa sslmode=disable"

	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Printf("Postgres not yet ready... %s", err)
			counts++
		} else {
			log.Println("Connected to Postgres!")
			return connection
		}

		if counts > 10 {
			log.Println(err)
			return nil
		}

		log.Println("Backing off for two seconds...")
		time.Sleep(2 * time.Second)
		continue
	}
}

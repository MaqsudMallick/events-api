package db

import (
	"database/sql"
	"fmt"
	"module-name/utils"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "postgres"
	dbname   = "eventdb"
)

var DB *sql.DB

func Init() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	var err error
	DB, err = sql.Open("postgres", psqlconn)
	utils.PanicOnError(err)

	err = DB.Ping()
	utils.PanicOnError(err)
	_, err = DB.Exec(`CREATE TABLE  IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	username TEXT NOT NULL UNIQUE,
	password TEXT NOT NULL
	)`)
	utils.PanicOnError(err)
	_, err = DB.Exec(`CREATE TABLE IF NOT EXISTS events (
	id SERIAL PRIMARY KEY,
	name TEXT NOT NULL, 
	description TEXT NOT NULL, 
	location TEXT NOT NULL, 
	date_time TIMESTAMP NOT NULL, 
	user_id INTEGER,
	FOREIGN KEY (user_id) REFERENCES users(id)
	)`)
	utils.PanicOnError(err)
	_, err = DB.Exec(`CREATE TABLE IF NOT EXISTS registrations (
	id SERIAL PRIMARY KEY,
	event_id INTEGER NOT NULL,
	user_id INTEGER NOT NULL,
	FOREIGN KEY (event_id) REFERENCES events(id),
	FOREIGN KEY (user_id) REFERENCES users(id)
	)`)
	utils.PanicOnError(err)
	fmt.Println("DB Connected.")
}

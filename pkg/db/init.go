package db

import "log"

// list of setup db
func Setup() {
	setupDB()
	log.Println("[................DB_name start................]")
}

// list of closing db
func Close() {
	closeDB()
	log.Println("[................DB_name close................]")
}
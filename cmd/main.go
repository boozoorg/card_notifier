package main

import (
	"card_notifier/api"
	"card_notifier/config"
	"card_notifier/internal/scheduler"
	"card_notifier/pkg/db"
	"log"
)

func main() {
	defer deferFunc()
	config.Setup("config/conf.json")
	log.Println("[................service start................]")
	go scheduler.Setup()
	db.Setup()
	api.Router()
}

func deferFunc()  {
	db.Close()
	log.Println("[................service close................]")
}
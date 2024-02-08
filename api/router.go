package api

import (
	"card_notifier/api/handler"
	"card_notifier/config"
	"fmt"
	"log"
	"net/http"
)

func Router() {
	router := http.NewServeMux()
	router.HandleFunc("/", handler.SaveEvent)
	log.Println("service run on port :8080...")
	log.Println(fmt.Sprintf("%s:%d", config.Configs.App.ServerName, config.Configs.App.Port))
	http.ListenAndServe(fmt.Sprintf("%s:%d", config.Configs.App.ServerName, config.Configs.App.Port), router)
}

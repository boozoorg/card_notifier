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
	log.Printf("[..........service run on port :%d..........]", config.Configs.App.Port)
	http.ListenAndServe(fmt.Sprintf("%s:%d", config.Configs.App.ServerName, config.Configs.App.Port), router)
}

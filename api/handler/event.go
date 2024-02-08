package handler

import (
	"card_notifier/internal/entities/event"
	"card_notifier/internal/utils/binder"
	"card_notifier/internal/utils/response"
	"log"
	"net/http"
)

func SaveEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		response.Message(http.StatusMethodNotAllowed, nil).ToClient(w)
		return
	}

	var input event.EventModel
	if err := binder.Bind(r, &input); err != nil {
		log.Println("binding err:", err)
		response.Message(http.StatusBadRequest, err).ToClient(w)
		return
	}

	resp, status, err := event.SaveEvent(input)
	if err != nil {
		response.Message(status, err).ToClient(w)
		return
	}

	response.Message(status, resp).ToClient(w)
}

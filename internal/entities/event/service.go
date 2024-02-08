package event

import (
	"card_notifier/internal/utils/response"
	"card_notifier/internal/utils/validation"
	"log"
	"net/http"
)

func SaveEvent(input EventModel) (resp interface{}, status int, err error) {
	if !validation.IsCardCorrect(input.Card) {
		log.Printf("card: '%v', is invalid", input.Card)
		return nil, http.StatusBadRequest, response.WrongCard
	}

	if !validation.IsWebUrlCorrect(input.WebsiteUrl) {
		log.Printf("web-url: '%v', is invalid", input.WebsiteUrl)
		return nil, http.StatusBadRequest, response.WrongWebUrl
	}

	if !validation.IsOrderTypeCorrect(input.OrderType) {
		log.Printf("order-type: '%v', is invalid", input.OrderType)
		return nil, http.StatusBadRequest, response.WrongWebUrl
	}

	return input.SaveInDB()
}

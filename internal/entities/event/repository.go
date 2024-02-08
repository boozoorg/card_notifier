package event

import (
	"card_notifier/pkg/db"
	"fmt"
	"net/http"
)

func (input EventModel) SaveInDB() (resp interface{}, status int, err error) {
	db.GetDB().Create(fmt.Sprintf(`{"orderType":"%v", "card":"%v", "eventDate":"%v", "websiteUrl":"%v"}`, input.OrderType, input.Card, input.EventDate, input.WebsiteUrl), input.SessionId)
	return `success`, http.StatusCreated, nil
}

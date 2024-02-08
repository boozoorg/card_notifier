package test

import (
	utils "card_notifier/internal/utils/validation"
	"testing"
)

func TestValidateOrderType(t *testing.T) {
	tests := []struct {
		name       string
		order      string
		wantResult bool
	}{
		// ToDo: add more test cases
		{name: "Correct Data", order: "CardVerify", wantResult: true},
		{name: "Correct Data", order: "Purchase", wantResult: true},
		{name: "Wrong Data", order: "SendOtd", wantResult: true},
		{name: "Empty Data", order: "", wantResult: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if utils.IsCardCorrect(tt.order) != tt.wantResult {
				t.Errorf("Card is invalid, input data: '%v'.", tt.order)
			}
		})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if utils.IsOrderTypeCorrect(tt.order) != tt.wantResult {
				t.Errorf("data is invalid, data: '%v'", tt.order)
			}
		})
	}
}

func TestValidateCard(t *testing.T) {
	tests := []struct {
		name       string
		card       string
		wantResult bool
	}{
		// ToDo: add more test cases
		{name: "Correct Data", card: "4433**1409", wantResult: true},
		{name: "Empty Data", card: "", wantResult: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if utils.IsCardCorrect(tt.card) != tt.wantResult {
				t.Errorf("data is invalid, input data: '%v'.", tt.card)
			}
		})
	}
}

func TestValidateWebUrl(t *testing.T) {
	tests := []struct {
		name       string
		page       string
		wantResult bool
	}{
		// ToDo: add more test cases
		{name: "Correct Data", page: "https://amazon.com", wantResult: true},
		{name: "Correct Data", page: "https://adidas.com", wantResult: true},
		{name: "Wrong Data", page: "https://somon.tj", wantResult: true},
		{name: "Empty Data", page: "", wantResult: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if utils.IsWebUrlCorrect(tt.page) != tt.wantResult {
				t.Errorf("data is invalid, data: '%v'", tt.page)
			}
		})
	}
}

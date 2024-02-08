package binder

import (
	"card_notifier/internal/utils/response"
	"encoding/json"
	"net/http"
)

// a simple bind function
func Bind(r *http.Request, input any) error {
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		if err.Error() == "EOF" { // if body is nil it will return EOF
			return response.EmptyBody
		}
		return err
	}
	return nil
}

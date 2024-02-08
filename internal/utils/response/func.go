package response

import (
	"encoding/json"
	"net/http"
)

type Msg struct {
	Status    int         // http status code
	ErrorCode knownErrors // is known error
	Body      Info        // what will be sent to client
}

type Info struct {
	Message interface{} `json:"message"` // response message
	Show    bool        `json:"show"`    // is show to client
}

// all service response handler
func Message(status int, resp interface{}) (mes Msg) {
	mes.Status = status
	// check is known error
	if v, ok := resp.(knownErrors); ok {
		mes.ErrorCode = v
		mes.Body.Show = true
		return
	}

	// check is error
	if v, ok := resp.(error); ok {
		mes.Body.Message = v.Error()
	} else {
		mes.Body.Message = resp
	}

	return
}

// handed response sender
func (mes Msg) ToClient(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	if mes.Body.Show {
		mes.Body.Message = err[mes.ErrorCode]
	}
	resp, _ := json.Marshal(mes.Body)
	w.WriteHeader(mes.Status)
	w.Write(resp)
}

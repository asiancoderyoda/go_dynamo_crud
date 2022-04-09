package http

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

func (res *Response) bytes() []byte {
	body, err := json.Marshal(res)
	if err != nil {
		log.Println(err)
	}
	return body
}

func (res *Response) string() string {
	body, err := json.Marshal(res)
	if err != nil {
		log.Println(err)
	}
	return string(body)
}

func (res *Response) response(w http.ResponseWriter, status int, body []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(body)
}

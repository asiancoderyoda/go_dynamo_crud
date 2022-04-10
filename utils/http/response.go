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
	return string(res.bytes())
}

func createResponse(data interface{}, status int) *Response {
	return &Response{
		Status: status,
		Data:   data,
	}
}

func (res *Response) response(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(res.Status)
	_, _ = w.Write(res.bytes())
	log.Println(res.string())
}

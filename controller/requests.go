package controller

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// добавил заполнение респонса исключительно для своего удобсва, по хорошему надо возвращать бы мапу стринг интерфейс
func MakeRequest(url string, response interface{}) {
	httpResponse, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	decodeResponse(httpResponse.Body, &response)
	defer httpResponse.Body.Close()
}

func decodeResponse(body io.Reader, response interface{}) {
	decoder := json.NewDecoder(body)
	err := decoder.Decode(&response)
	if err != nil {
		log.Fatal(err)
	}
}

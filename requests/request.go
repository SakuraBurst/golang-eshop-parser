package requests

import (
	"log"
	"net/http"
)

func MakeRequest(url string, response interface{}) {
	httpResponse, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	decodeResponse(httpResponse.Body, &response)
	defer httpResponse.Body.Close()
}

package requests

import (
	"eshop-parser/utils"
	"log"
	"net/http"
)

func MakeRequest(url string, response interface{}) {
	httpResponse, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	utils.FillJson(httpResponse.Body, &response)
	defer httpResponse.Body.Close()
}

package controller

import (
	"encoding/json"
	"eshop-parser/workers"
	"fmt"
	"net/http"
)

var GamePricesService http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
	encoder := json.NewEncoder(writer)
	if request.Method != http.MethodPost {
		return
	}
	fmt.Println(request.Body)
	respose := workers.GetGamePricesFromJson(request.Body)
	encoder.Encode(respose)
	return
}

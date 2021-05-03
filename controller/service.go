package controller

import (
	"encoding/json"
	"eshop-parser/workers"
	"net/http"
)

var GamePricesService http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
	encoder := json.NewEncoder(writer)
	if request.Method != http.MethodPost {
		return
	}
	response := workers.GetGamePricesFromJson(request.Body)
	encoder.Encode(response)
	return
}

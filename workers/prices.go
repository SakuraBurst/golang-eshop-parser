package workers

import (
	"eshop-parser/environment"
	"eshop-parser/models"
	"fmt"
	"io"
)

func GetGamePricesFromJson(body io.ReadCloser) []models.GameResponse {
	requestSlice := environment.CreateSliceOfGameRequestFromJson(body, "games.json")
	var responseSlice []models.GameResponse
	for i, v := range requestSlice {
		fmt.Println(i)
		go requestSlice[i].Request()
		response := models.GameResponse{GameName: v.GameName, GameInfo: make(map[string]interface{})}
		responseSlice = append(responseSlice, response)
	}
	for i := range responseSlice {
		fmt.Println(i)
		responseSlice[i].GameInfo = <-requestSlice[i].ResponseChannel
	}
	fmt.Println(responseSlice)
	return responseSlice
}

package main

import (
	"eshop-parser/environment"
	"eshop-parser/models"
	"fmt"
)

func main() {
	requestSlice := environment.CreateSliceOfGameRequestFromJson("games.json")
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
}

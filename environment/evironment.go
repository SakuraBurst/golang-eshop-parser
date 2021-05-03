package environment

import (
	"eshop-parser/models"
	"eshop-parser/repository"
	"eshop-parser/utils"
	"io"
)

func CreateSliceOfGameRequestFromJson(body io.ReadCloser, jsonFileName string) []models.GameRequest {
	gameJsonFile := repository.OpenFile(jsonFileName)
	gameSlice := getGameSlice(body, gameJsonFile)
	var requestSlice []models.GameRequest
	for _, value := range gameSlice {
		request := models.GameRequest{GameName: value["name"], GameId: value["id"], ResponseChannel: make(chan map[string]interface{})}
		requestSlice = append(requestSlice, request)
	}
	return requestSlice
}

func getGameSlice(reader io.Reader, writer io.WriterAt) models.GamesSlice {
	gameSlice := models.GamesSlice{}
	utils.FillJson(reader, &gameSlice)
	searchForGameIds(gameSlice, writer)
	return gameSlice
}

func searchForGameIds(gamesSlice models.GamesSlice, jsonFile io.WriterAt) models.GamesSlice {
	gamesSlice.GetGameIds()
	repository.WriteFile(jsonFile, gamesSlice)
	return gamesSlice
}

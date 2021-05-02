package environment

import (
	"eshop-parser/models"
	"eshop-parser/repository"
	"fmt"
	"io"
)

func getGameMap(reader io.Reader, writer io.WriterAt) models.GamesSlice {
	gameJsonMap := models.GamesSlice{}
	repository.FillJsonMap(reader, &gameJsonMap)
	checkIsAllGameIdIsExist(gameJsonMap, writer)
	fmt.Println(gameJsonMap)
	return gameJsonMap
}

func checkIsAllGameIdIsExist(gamesMap models.GamesSlice, jsonFile io.WriterAt) models.GamesSlice {

	gamesMap.GetGameIds()
	repository.WriteFile(jsonFile, gamesMap)

	return gamesMap
}

func CreateSliceOfGameRequestFromJson(body io.ReadCloser, jsonFileName string) []models.GameRequest {
	gameJsonFile := repository.OpenFile(jsonFileName)
	gameMap := getGameMap(body, gameJsonFile)
	var requestSlice []models.GameRequest
	for _, value := range gameMap {
		request := models.GameRequest{GameName: value["name"], GameId: value["id"], ResponseChannel: make(chan map[string]interface{})}
		requestSlice = append(requestSlice, request)
	}
	return requestSlice
}

package environment

import (
	"eshop-parser/models"
	"eshop-parser/repository"
	"fmt"
	"io"
)

func getGameMap(reader io.Reader, writer io.WriterAt) models.GamesMap {
	gameJsonMap := models.GamesMap{}
	repository.FillJsonMap(reader, &gameJsonMap)
	checkIsAllGameIdIsExist(gameJsonMap, writer)
	fmt.Println(gameJsonMap)
	return gameJsonMap
}

func checkIsAllGameIdIsExist(gamesMap models.GamesMap, jsonFile io.WriterAt) models.GamesMap {
	if !gamesMap.IsAllGamesHasId() {
		gamesMap.GetGameIds()
		repository.WriteFile(jsonFile, gamesMap)
	}
	return gamesMap
}

func CreateSliceOfGameRequestFromJsonFile(jsonFileName string) []models.GameRequest {
	gameJsonFile := repository.OpenFile(jsonFileName)
	gameMap := getGameMap(gameJsonFile, gameJsonFile)
	var requestSlice []models.GameRequest
	for key, value := range gameMap {
		request := models.GameRequest{GameName: key, GameId: value["id"], ResponseChannel: make(chan map[string]interface{})}
		requestSlice = append(requestSlice, request)
	}
	return requestSlice
}

func CreateSliceOfGameRequestFromJson(body io.ReadCloser, jsonFileName string) []models.GameRequest {
	gameJsonFile := repository.OpenFile(jsonFileName)
	gameMap := getGameMap(body, gameJsonFile)
	var requestSlice []models.GameRequest
	for key, value := range gameMap {
		request := models.GameRequest{GameName: key, GameId: value["id"], ResponseChannel: make(chan map[string]interface{})}
		requestSlice = append(requestSlice, request)
	}
	return requestSlice
}

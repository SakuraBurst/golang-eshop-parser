package environment

import (
	"eshop-parser/controller"
	"eshop-parser/models"
	"eshop-parser/repository"
	"fmt"
	"io"
)

func getGameMap(filename string) models.GamesMap {
	gameJsonFile := repository.OpenFile(filename)
	gameJsonMap := models.GamesMap{}
	controller.FillFilesMap(gameJsonFile, &gameJsonMap)
	checkIsAllGameIdIsExist(gameJsonMap, gameJsonFile)
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

func CreateSliceOfGameRequestFromJson(jsonName string) []models.GameRequest {
	gameMap := getGameMap(jsonName)
	var requestSlice []models.GameRequest
	for key, value := range gameMap {
		request := models.GameRequest{GameName: key, GameId: value["id"], ResponseChannel: make(chan map[string]interface{})}
		requestSlice = append(requestSlice, request)
	}
	return requestSlice
}

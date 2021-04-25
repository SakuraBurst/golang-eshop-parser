package environment

import (
	"encoding/json"
	"eshop-parser/models"
	"eshop-parser/repository"
	"fmt"
	"io"
	"log"
)

func GetGameIds(filename string) []string {
	gameJsonFile := repository.OpenFile(filename)
	gameJsonMap := getGamesMap(gameJsonFile)
	checkIsAllGameIdIsExist(gameJsonMap, gameJsonFile)
	fmt.Println(gameJsonMap)
	return []string{"dfd"}
}

func getGamesMap(jsonFile io.Reader) models.GamesMap {
	var file models.GamesMap
	decoder := json.NewDecoder(jsonFile)
	err := decoder.Decode(&file)
	if err != nil {
		log.Fatal("ошибочка при чтении файла", err)
	}
	return file
}

func checkIsAllGameIdIsExist(gamesMap models.GamesMap, jsonFile io.WriterAt) models.GamesMap {
	if !gamesMap.IsAllGamesHasId() {
		gamesMap.GetGameIds()
		repository.WriteFile(jsonFile, gamesMap)
	}
	return gamesMap
}
func returnGameIds() {

}

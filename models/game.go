package models

import (
	"eshop-parser/searcher"
)

type GamesSlice []GameFromJson

type GameFromJson map[string]string

func (gamesSlice GamesSlice) GetGameIds() GamesSlice {
	requestChannels := make([]chan string, 0)
	for ind, value := range gamesSlice {
		requestChannels = append(requestChannels, make(chan string))
		go searcher.Searcher(value["name"], requestChannels[ind])
	}
	for index, ch := range requestChannels {
		gamesSlice[index]["id"] = <-ch
	}
	return gamesSlice
}

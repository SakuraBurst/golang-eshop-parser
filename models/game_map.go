package models

import (
	searcher3 "eshop-parser/searcher"
)

type GamesSlice []GameFromJson

func (gamesMap GamesSlice) GetGameIds() GamesSlice {
	requestChannels := make([]chan string, 0)
	for ind, value := range gamesMap {
		requestChannels = append(requestChannels, make(chan string))
		go searcher3.Searcher(value["name"], requestChannels[ind])
	}
	for index, ch := range requestChannels {
		gamesMap[index]["id"] = <-ch
	}
	return gamesMap
}

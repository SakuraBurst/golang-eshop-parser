package models

import "fmt"

type GamesSlice []GameFromJson

func (gamesMap GamesSlice) GetGameIds() GamesSlice {
	requestChannels := make([]chan string, 0)
	for ind, value := range gamesMap {
		fmt.Println(ind)
		requestChannels = append(requestChannels, make(chan string))
		go searcher(value["name"], requestChannels[ind])
	}
	for index, ch := range requestChannels {
		fmt.Println(index)
		gamesMap[index]["id"] = <-ch
	}
	return gamesMap
}

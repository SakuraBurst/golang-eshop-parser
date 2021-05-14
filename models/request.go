package models

import (
	"eshop-parser/searcher"
)

type Requester interface {
	GetIds() []Game
	GetPrices() []Game
	GetGameSlice() []Game
}

type EshopGameRequester struct {
	GameSlice []Game
	Searcher  searcher.Searcher
}

func (gReq *EshopGameRequester) GetIds() []Game {
	requestChannels := make([]chan string, 0)
	for ind, value := range gReq.GameSlice {
		requestChannels = append(requestChannels, make(chan string))
		go gReq.Searcher.SearchForId(value.GameName, requestChannels[ind])
	}
	for index, ch := range requestChannels {
		gReq.GameSlice[index].GameId = <-ch
	}
	gReq.GameSlice = clearEmptyIds(gReq.GameSlice)
	return gReq.GameSlice
}

func (gReq EshopGameRequester) GetPrices() []Game {
	requestChannels := make([]chan map[string]interface{}, 0)
	for ind := range gReq.GameSlice {
		requestChannels = append(requestChannels, make(chan map[string]interface{}))
		go gReq.Searcher.SearchForPrice(gReq.GameSlice[ind].GameId, requestChannels[ind])
	}
	for index, ch := range requestChannels {
		gReq.GameSlice[index].GameInfo = <-ch
	}
	gReq.GameSlice = clearEmptyGameInfos(gReq.GameSlice)
	return gReq.GameSlice
}

func (gReq EshopGameRequester) GetGameSlice() []Game {
	return gReq.GameSlice
}

func clearEmptyIds(games []Game) []Game {
	var filtredGames []Game
	for i := 0; i < len(games); i++ {
		if games[i].GameId != "" {
			filtredGames = append(filtredGames, games[i])
		}
	}
	return filtredGames
}

func clearEmptyGameInfos(games []Game) []Game {
	var filtredGames []Game
	for i := 0; i < len(games); i++ {
		if games[i].GameInfo != nil {
			filtredGames = append(filtredGames, games[i])
		}
	}
	return filtredGames
}

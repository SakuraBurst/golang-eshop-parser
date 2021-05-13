package eshop

import (
	"eshop-parser/requests"
	"eshop-parser/searcher"
	"eshop-parser/utils"
	"fmt"
	"net/url"
	"regexp"
)

func (s EshopSearcher) SearchForId(gameName string, idChannel chan string) {
	searchUrl := createUrl(gameName)
	response, err := makeRequest(searchUrl)
	if err != nil {
		idChannel <- ""
		return
	}
	typedGamesSlice := getTypedGamesSlice(response)
	game := getGame(typedGamesSlice, gameName)
	gameId := getGameId(game)
	idChannel <- gameId
}

func createUrl(gameName string) string {
	originalUrl := "https://searching.nintendo-europe.com/ru/select?"
	params := url.Values{}
	params.Add("fq", "type:GAME AND ((playable_on_txt:\"HAC\"))")
	params.Add("q", gameName)
	return originalUrl + params.Encode()
}

func makeRequest(url string) (searcher.Response, error) {
	var response searcher.Response
	err := requests.MakeRequest(url, &response)
	return response, err
}

func getTypedGamesSlice(resp searcher.Response) searcher.ResponseGameSlice {
	untypedGamesSlice := resp["response"]["docs"]
	var typedGamesSlice searcher.ResponseGameSlice
	utils.ReDecodeToNewJson(untypedGamesSlice, &typedGamesSlice)
	return typedGamesSlice
}

func getGame(gamesSlice searcher.ResponseGameSlice, gameName string) searcher.ResponseGame {
	//Loop:
	for i, v := range gamesSlice {
		title := v["title"]
		switch titleString := title.(type) {
		case string:
			matched, _ := regexp.MatchString(fmt.Sprintf("(?i)^%v", gameName), titleString)
			if matched {
				switch gamesSlice[i]["nsuid_txt"].(type) {
				case []interface{}:
					return gamesSlice[i]
				}
				//break Loop
			}
		}
	}
	return searcher.ResponseGame{}
}

func getGameId(game searcher.ResponseGame) string {
	switch idSlice := game["nsuid_txt"].(type) {
	case []interface{}:
		return idSlice[0].(string)
	}
	return ""
}

package searcher

import (
	"eshop-parser/requests"
	"eshop-parser/utils"
	"fmt"
	"net/url"
	"regexp"
)

func Searcher(gameName string, idChannel chan string) {
	searchUrl := createUrl(gameName)
	response := makeRequest(searchUrl)
	typedGamesSlice := getTypedGamesSlice(response)
	game := getGame(typedGamesSlice, gameName)
	gameId := getGameId(game)
	idChannel <- gameId
}

func makeRequest(url string) Response {
	var response Response
	requests.MakeRequest(url, &response)
	return response
}

func getTypedGamesSlice(resp Response) ResponseGameSlice {
	untypedGamesSlice := resp["response"]["docs"]
	var typedGamesSlice ResponseGameSlice
	utils.ReDecodeToNewJson(untypedGamesSlice, &typedGamesSlice)
	return typedGamesSlice
}

func createUrl(gameName string) string {
	var originalUrl string = "https://searching.nintendo-europe.com/ru/select?"
	params := url.Values{}
	params.Add("fq", "type:GAME AND ((playable_on_txt:\"HAC\"))")
	params.Add("q", gameName)
	return originalUrl + params.Encode()
}

func getGame(gamesSlice ResponseGameSlice, gameName string) ResponseGame {
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
	return ResponseGame{}
}

func getGameId(game ResponseGame) string {
	switch idSlice := game["nsuid_txt"].(type) {
	case []interface{}:
		return idSlice[0].(string)
	}
	return ""
}

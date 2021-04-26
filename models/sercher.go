package models

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"regexp"
)

func searcher(gameName string) string {
	response := makeRequest(gameName)
	typedGamesSlice := getTypedGamesSlice(response)
	game := getGame(typedGamesSlice, gameName)
	gameId := getGameId(game)
	fmt.Println(game["title"])
	fmt.Println(gameId)
	return gameId
}

func getTypedGamesSlice(resp Response) ResponseGameSlice {
	untypedGamesSlice := resp["response"]["docs"]
	var semiTypedGamesSlice []interface{}
	var typedGamesSlice ResponseGameSlice
	switch typedSlice := untypedGamesSlice.(type) {
	case []interface{}:
		semiTypedGamesSlice = typedSlice
		for _, v := range semiTypedGamesSlice {
			switch typedMap := v.(type) {
			case map[string]interface{}:
				typedGamesSlice = append(typedGamesSlice, typedMap)
			}
		}
	}
	return typedGamesSlice
}

func makeRequest(gameName string) Response {
	searchUrl := createUrl(gameName)
	httpResponse, err := http.Get(searchUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer httpResponse.Body.Close()
	response := decodeResponse(httpResponse.Body)
	return response
}

func createUrl(gameName string) string {
	var originalUrl string = "https://searching.nintendo-europe.com/ru/select?"
	params := url.Values{}
	params.Add("fq", "type:GAME AND ((playable_on_txt:\"HAC\"))")
	params.Add("q", gameName)
	return originalUrl + params.Encode()
}

func decodeResponse(body io.Reader) Response {
	var response Response
	decoder := json.NewDecoder(body)
	err := decoder.Decode(&response)
	if err != nil {
		log.Fatal(err)
	}
	return response
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
	//getGame(response, gameName)
	return ""
}

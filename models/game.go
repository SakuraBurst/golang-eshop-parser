package models

import (
	"eshop-parser/requests"
)

type GameRequest struct {
	ResponseChannel chan map[string]interface{}
	GameName        string
	GameId          string
}

type GameResponse struct {
	GameName string                 `json:"game_name"`
	GameInfo map[string]interface{} `json:"game_info"`
}

func (gReq *GameRequest) Request() {
	var resp map[string]interface{}
	requests.MakeRequest("https://api.ec.nintendo.com/v1/price?country=RU&lang=ru&ids="+gReq.GameId, &resp)
	gReq.ResponseChannel <- resp
}

type GameFromJson map[string]string

func (game GameFromJson) isIdExist() bool {
	return len(game["id"]) > 0
}

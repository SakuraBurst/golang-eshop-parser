package models

type Game struct {
	GameName string `json:"name"`
	GameId   string
	GameInfo map[string]interface{} `json:"game_info"`
}

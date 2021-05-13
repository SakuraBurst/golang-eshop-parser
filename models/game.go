package models

type Game struct {
	GameName string                 `json:"name"`
	GameId   string                 `json:"game_id"`
	GameInfo map[string]interface{} `json:"game_info"`
}

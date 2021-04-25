package models

type Game map[string]string

func (game Game) isIdExist() bool {
	return len(game["id"]) > 0
}

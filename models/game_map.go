package models

type GamesMap map[string]GameFromJson

func (gamesMap GamesMap) GetGameIds() GamesMap {
	for key, val := range gamesMap {
		if !val.isIdExist() {
			val["id"] = searcher(key)
		}
	}
	return gamesMap
}

func (gamesMap GamesMap) IsAllGamesHasId() bool {
	for _, val := range gamesMap {
		if !val.isIdExist() {
			return false
		}

	}
	return true
}

package environment

import (
	"eshop-parser/models"
	"eshop-parser/utils"
	"io"
)

func CreateSliceOfEshopGameRequestFromJson(body io.ReadCloser) models.Requester {
	gameSlice := getGameSlice(body)
	gameRequester := models.EshopGameRequester{GameSlice: gameSlice}
	return &gameRequester
}

func getGameSlice(reader io.Reader) []models.Game {
	gameSlice := make([]models.Game, 0)
	utils.FillJson(reader, &gameSlice)
	return gameSlice
}

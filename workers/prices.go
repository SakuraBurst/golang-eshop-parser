package workers

import (
	"eshop-parser/environment"
	"eshop-parser/models"
	"io"
)

func GetGamePricesFromJson(body io.Reader) []models.Game {
	gameRequester := environment.CreateEshopGameRequesterFromJson(body)
	gameRequester.GetIds()
	return gameRequester.GetPrices()
}

package workers

import (
	"eshop-parser/environment"
	"eshop-parser/models"
	"io"
)

func GetGamePricesFromJson(body io.ReadCloser) []models.Game {
	gameRequester := environment.CreateSliceOfEshopGameRequestFromJson(body)
	gameRequester.GetIds()
	return gameRequester.GetPrices()
}

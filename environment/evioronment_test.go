package environment_test

import (
	"eshop-parser/environment"
	"eshop-parser/models"
	"strings"
	"testing"
)

var mock = []models.Game{{
	GameName: "Minecraft",
}}

func TestEshopSliceCreator(t *testing.T) {
	testBody := strings.NewReader(`[{"name": "Minecraft"}]`)
	requester := environment.CreateEshopGameRequesterFromJson(testBody).GetGameSlice()
	if len(requester) != len(mock) {
		t.Errorf("wrong game length")
	}
	if requester[0].GameName != mock[0].GameName {
		t.Errorf("something wrong with puted game")
	}
}

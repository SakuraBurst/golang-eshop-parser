package utils_test

import (
	"eshop-parser/utils"
	"strings"
	"testing"
)

func TestMapperFiller(t *testing.T) {
	testReader := strings.NewReader(`[{"name": "Minecraft"}]`)
	testTarget := make([]map[string]string, 0)
	utils.FillJson(testReader, &testTarget)
	if testTarget[0]["name"] != "Minecraft" {
		t.Errorf("Filler Error")
	}
}

func TestMapperReDecoder(t *testing.T) {
	testMap := make([]map[string]interface{}, 1)
	testMap[0] = map[string]interface{}{"name": "Minecraft"}
	testTarget := make([]map[string]string, 0)
	utils.ReDecodeToNewJson(testMap, &testTarget)
	if testTarget[0]["name"] != "Minecraft" {
		t.Errorf("Filler Error")
	}
}

package utils

import (
	"encoding/json"
	"io"
)

func FillJson(reader io.Reader, data interface{}) {
	decoder := json.NewDecoder(reader)
	err := decoder.Decode(&data)
	CheckErrorAndShotDownIfItIs(err, "")
}

func ReDecodeToNewJson(source interface{}, data interface{}) {
	newJson, err := json.Marshal(source)
	CheckErrorAndShotDownIfItIs(err, "прикол в редекоде")
	err = json.Unmarshal(newJson, &data)
	CheckErrorAndShotDownIfItIs(err, "прикол в редекоде")
}

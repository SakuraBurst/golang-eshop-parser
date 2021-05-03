package utils

import (
	"encoding/json"
	"io"
	"log"
)

func FillJson(reader io.Reader, data interface{}) {
	decoder := json.NewDecoder(reader)
	err := decoder.Decode(&data)
	if err != nil {
		log.Fatal(err)
	}
}

func ReDecodeToNewJson(source interface{}, data interface{}) {
	newJson, err := json.Marshal(source)
	if err != nil {
		log.Fatal("прикол в редекоде", err)
	}
	err = json.Unmarshal(newJson, &data)
	if err != nil {
		log.Fatal("прикол в редекоде", err)
	}
}

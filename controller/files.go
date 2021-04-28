package controller

import (
	"encoding/json"
	"io"
	"log"
)

func FillFilesMap(jsonFile io.Reader, file interface{}) {
	decoder := json.NewDecoder(jsonFile)
	err := decoder.Decode(&file)
	if err != nil {
		log.Fatal("ошибочка при чтении файла", err)
	}
}

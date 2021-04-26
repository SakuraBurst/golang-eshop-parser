package controller

import (
	"encoding/json"
	"eshop-parser/models"
	"io"
	"log"
)

func FillFilesMap(jsonFile io.Reader, file models.ParsedFileInterface) {
	decoder := json.NewDecoder(jsonFile)
	err := decoder.Decode(&file)
	if err != nil {
		log.Fatal("ошибочка при чтении файла", err)
	}
}

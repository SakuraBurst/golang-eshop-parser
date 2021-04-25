package repository

import (
	"encoding/json"
	"eshop-parser/models"
	"io"
	"log"
	"os"
)

func OpenFile(filename string) *os.File {
	f, err := os.OpenFile(filename, os.O_RDWR, 0777)
	if err != nil {
		log.Fatal("ошибочка при открытии файла: ", err)
	}
	return f
}

func WriteFile(file io.WriterAt, gamesMap models.GamesMap) {
	newJson, err := json.MarshalIndent(gamesMap, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	_, err = file.WriteAt(newJson, 0)
	if err != nil {
		log.Fatal(err)
	}
}

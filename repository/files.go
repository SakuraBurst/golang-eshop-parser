package repository

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

func OpenFile(filename string) *os.File {
	//f, err := os.OpenFile(filename, os.O_RDWR, 0777)
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal("ошибочка при открытии файла: ", err)
	}
	return f
}

func WriteFile(file io.WriterAt, content interface{}) {
	switch typedContent := content.(type) {
	case []byte:
		_, err := file.WriteAt(typedContent, 0)
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	newJson, err := json.MarshalIndent(content, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	_, err = file.WriteAt(newJson, 0)
	if err != nil {
		log.Fatal(err)
	}
}

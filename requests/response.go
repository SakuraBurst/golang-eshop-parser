package requests

import (
	"encoding/json"
	"io"
	"log"
)

func decodeResponse(body io.Reader, response interface{}) {
	decoder := json.NewDecoder(body)
	err := decoder.Decode(&response)
	if err != nil {
		log.Fatal(err)
	}
}

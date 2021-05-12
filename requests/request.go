package requests

import (
	"eshop-parser/utils"
	"net/http"
)

func MakeRequest(url string, response interface{}) error {
	httpResponse, err := http.Get(url)
	if err != nil {
		return err
	}
	utils.FillJson(httpResponse.Body, &response)
	defer httpResponse.Body.Close()
	return nil
}

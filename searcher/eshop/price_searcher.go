package eshop

import "eshop-parser/requests"

func SearchForPrice(id string, responseChannel chan map[string]interface{}) {
	var resp map[string]interface{}
	err := requests.MakeRequest("https://api.ec.nintendo.com/v1/price?country=RU&lang=ru&ids="+id, &resp)
	if err != nil {
		responseChannel <- nil
		return
	}
	responseChannel <- resp
}

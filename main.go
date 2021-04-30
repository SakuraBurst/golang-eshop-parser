package main

import (
	"eshop-parser/controller"
	"log"
	"net/http"
)

func main() {
	http.Handle("/get", controller.GamePricesService)
	log.Fatal(http.ListenAndServe(":8181", nil))
}

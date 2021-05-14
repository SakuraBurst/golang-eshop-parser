package eshop_test

import (
	"eshop-parser/searcher/eshop"
	"fmt"
	"testing"
)

var MinecraftId = "70010000000963"

func TestEshopSearcher_SearchForId(t *testing.T) {
	searcher := eshop.EshopSearcher{}
	idChannel := make(chan string)
	go searcher.SearchForId("Minecraft", idChannel)
	result := <-idChannel
	fmt.Println(result)
	if len(result) == 0 {
		t.Errorf("No result")
	}
	if result != MinecraftId {
		t.Errorf("Wrong result")
	}
}

func TestEshopSearcher_SearchForPrice(t *testing.T) {
	searcher := eshop.EshopSearcher{}
	infoChannel := make(chan map[string]interface{})
	go searcher.SearchForPrice(MinecraftId, infoChannel)
	result := <-infoChannel
	if result == nil {
		t.Errorf("No result")
	}
}

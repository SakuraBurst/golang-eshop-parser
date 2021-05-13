package searcher

type Searcher interface {
	SearchForId(gameName string, idChannel chan string)
	SearchForPrice(id string, responseChannel chan map[string]interface{})
}

type Response map[string]map[string]interface{}

type ResponseGameSlice []ResponseGame

type ResponseGame map[string]interface{}

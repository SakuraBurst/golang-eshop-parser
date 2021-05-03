package models

type RequestMaker interface {
	MakeRequest(url string, resp interface{})
}

type Request struct {
	ResponseChannel chan interface{}
	Url             string
	RequestMaker    RequestMaker
	Resp            interface{}
}

func (aReq *Request) Request() {
	aReq.RequestMaker.MakeRequest(aReq.Url, aReq.Resp)
	aReq.ResponseChannel <- aReq.Resp
}

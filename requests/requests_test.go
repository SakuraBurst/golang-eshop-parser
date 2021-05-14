package requests_test

import (
	"eshop-parser/requests"
	"testing"
)

func TestRequest(t *testing.T) {
	response := make(map[string]interface{})
	err := requests.MakeRequest("https://jsonplaceholder.typicode.com/todos/1", &response)
	if err != nil {
		t.Errorf("Request error")
	}
	if len(response) == 0 {
		t.Errorf("Response error")
	}
}

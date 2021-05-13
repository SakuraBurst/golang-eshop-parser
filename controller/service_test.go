package controller_test

import (
	"eshop-parser/controller"
	"fmt"
	"io"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"
)

type TestRequest struct {
	testName string
	method   string
	target   string
	body     io.Reader
	wantBody *regexp.Regexp
}

var TestRequests = []TestRequest{
	{
		testName: "Гет запрос",
		method:   "GET",
		target:   "/",
		body:     nil,
		wantBody: nil,
	}, {
		testName: "Пост запрос с одной игрой",
		method:   "POST",
		target:   "/",
		body:     strings.NewReader(`[{"name": "Minecraft"}]`),
		wantBody: regexp.MustCompile(`\[{"name":"Minecraft","game_id":"\d+","game_info":[\w\W]+`),
	}}

func TestService(t *testing.T) {
	for _, tt := range TestRequests {
		t.Run(tt.testName, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, tt.target, tt.body)
			w := httptest.NewRecorder()
			controller.GamePricesService(w, req)
			resp := w.Result()
			body, _ := io.ReadAll(resp.Body)
			if tt.wantBody == nil && len(string(body)) > 0 {
				t.Errorf("expected no body")
			}
			if tt.wantBody != nil && !tt.wantBody.MatchString(string(body)) {
				t.Errorf("wrong body")
			}
			fmt.Println(string(body))
		})
	}
}

package rest_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mtolley/hello-api/handlers/rest"
)

func TestTranslateAPI(t *testing.T) {
	tt := []struct {
		Endpoint            string
		StatusCode          int
		ExpectedLanguage    string
		ExpectedTranslation string
	}{
		{
			Endpoint:            "/hello",
			StatusCode:          http.StatusOK,
			ExpectedLanguage:    "english",
			ExpectedTranslation: "hello",
		},
		{
			Endpoint:            "/hello?language=german",
			StatusCode:          http.StatusOK,
			ExpectedLanguage:    "german",
			ExpectedTranslation: "hallo",
		},
		{
			Endpoint:            "/hello?language=dutch",
			StatusCode:          http.StatusNotFound,
			ExpectedLanguage:    "",
			ExpectedTranslation: "",
		},
	}

	handler := http.HandlerFunc(rest.TranslateHandler)

	for _, test := range tt {
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", test.Endpoint, nil)

		handler.ServeHTTP(rr, req)

		if rr.Code != test.StatusCode {
			t.Errorf("Expected status %d but received %d", test.StatusCode, rr.Code)
		}

		var resp rest.Resp
		json.Unmarshal(rr.Body.Bytes(), &resp)

		if resp.Translation != test.ExpectedTranslation {
			t.Errorf("Expected translation to be %s but received %s", test.ExpectedTranslation, resp.Translation)
		}
	}
}

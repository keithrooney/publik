package internal

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDispatch(t *testing.T) {

	dispatcher := NewDispatcher()

	t.Run("TestDispatchReturns200", func(t *testing.T) {

		req, err := http.NewRequest("GET", "/v1/echo", nil)
		if err != nil {
			t.Fatal(err)
		}

		req.RemoteAddr = "165.45.67.202:422"

		response := httptest.NewRecorder()

		dispatcher.ServeHTTP(response, req)

		if response.Code != http.StatusOK {
			t.Errorf("expected code to %d, actual code is %d", http.StatusOK, response.Code)
		}

		expected := Response{
			IP: "165.45.67.202",
		}
		actual := &Response{}
		if err := json.Unmarshal(response.Body.Bytes(), actual); err != nil {
			t.Error("failed to read body")
		}

		if expected != *actual {
			t.Error("expected != actual")
		}

	})

	t.Run("TestDispatchReturns404ForUnsupportedMethod", func(t *testing.T) {

		req, err := http.NewRequest("POST", "/v1/echo", nil)
		if err != nil {
			t.Fatal(err)
		}
		response := httptest.NewRecorder()

		dispatcher.ServeHTTP(response, req)

		if response.Code != http.StatusBadRequest {
			t.Errorf("expected code to %d, actual code is %d", http.StatusBadRequest, response.Code)
		}

	})

	t.Run("TestDispatchReturns404ForInvalidUrl", func(t *testing.T) {

		req, err := http.NewRequest("POST", "/v1/echo", nil)
		if err != nil {
			t.Fatal(err)
		}
		response := httptest.NewRecorder()

		req.RemoteAddr = "165.45.67.202"

		dispatcher.ServeHTTP(response, req)

		if response.Code != http.StatusBadRequest {
			t.Errorf("expected code to %d, actual code is %d", http.StatusBadRequest, response.Code)
		}

	})

}

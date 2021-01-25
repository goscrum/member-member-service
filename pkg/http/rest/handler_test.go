package rest

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	handler := SetupHandler()

	if handler == nil {
		t.Error("Expected a http.SetupHandler instance, got nil")
	}
}

func Test_healthcheck(t *testing.T) {
	const expectedStatusCode = 200
	const expectedBody = `{"message":"Service is healthy","status":200}`

	handler := SetupHandler()
	r := httptest.NewRequest(http.MethodGet, "/api/v1/healthcheck", nil)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, r)

	if w.Result().StatusCode != expectedStatusCode {
		t.Errorf("Expected status code %v, got %v", expectedStatusCode, w.Result().StatusCode)
	}
	assertResponseJSON(t, w, expectedBody)
}

func assertResponseJSON(t *testing.T, w *httptest.ResponseRecorder, expectedResponseJSON string) {
	t.Helper()
	respBodyBytes, err := ioutil.ReadAll(w.Body)
	respBody := string(bytes.TrimSpace(respBodyBytes))
	if err != nil {
		t.Fatal("Unable to read response from Recorder")
	}
	if respBody != expectedResponseJSON {
		t.Errorf("Expected response %s; got %s", expectedResponseJSON, respBody)
	}
}

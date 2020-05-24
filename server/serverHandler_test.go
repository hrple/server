package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServerGet(t *testing.T) {
	var testServer = NewServer()
	err := testServer.Get("/test", &HrpleTestHandler{})

	if err != nil {
		t.Errorf("Error Occurred: %v", err)
	}

	request, _ := http.NewRequest(http.MethodGet, "/test", nil)
	response := httptest.NewRecorder()

	testServer.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("Something wrong, code : %v", response.Code)
	}
}

func TestServerPut(t *testing.T) {
	var testServer = NewServer()
	err := testServer.Put("/test", &HrpleTestHandler{})

	if err != nil {
		t.Errorf("Error Occurred: %v", err)
	}

	request, _ := http.NewRequest(http.MethodPut, "/test", nil)
	response := httptest.NewRecorder()

	testServer.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("Something wrong")
	}
}

func TestServerPutHandlerNotFound(t *testing.T) {
	var testServer = NewServer()
	err := testServer.Put("/test", &HrpleTestHandler{})

	if err != nil {
		t.Errorf("Error Occurred: %v", err)
	}

	request, _ := http.NewRequest(http.MethodPut, "/testFuncNotFound", nil)
	response := httptest.NewRecorder()

	testServer.ServeHTTP(response, request)

	if response.Code == http.StatusOK {
		t.Fatalf("Something wrong")
	}
}

func TestServerPutSubHandlerNotFound(t *testing.T) {
	var testServer = NewServer()
	err := testServer.Put("/mufasa", &HrpleTestHandler{})

	if err != nil {
		t.Errorf("Error Occurred: %v", err)
	}

	request, _ := http.NewRequest(http.MethodPut, "/testFuncNotFound/Fail", nil)
	response := httptest.NewRecorder()

	testServer.ServeHTTP(response, request)

	if response.Code == http.StatusOK {
		t.Fatalf("Something wrong")
	}
}

func TestServerPutMethodNotFound(t *testing.T) {
	var testServer = NewServer()
	err := testServer.Put("/test", &HrpleTestHandler{})

	if err != nil {
		t.Errorf("Error Occurred: %v", err)
	}

	request, _ := http.NewRequest(http.MethodGet, "/test", nil)
	response := httptest.NewRecorder()

	testServer.ServeHTTP(response, request)

	if response.Code == http.StatusOK {
		t.Fatalf("Something wrong")
	}
}

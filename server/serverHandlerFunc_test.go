package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServerGetFunc(t *testing.T) {
	var testServer = NewServer()
	err := testServer.GetFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

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

func TestServerPutFunc(t *testing.T) {
	var testServer = NewServer()
	err := testServer.PutFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

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

func TestServerPutFuncHandlerNotFound(t *testing.T) {
	var testServer = NewServer()
	err := testServer.PutFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		t.Log("PUT")
	})

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

func TestServerPutFuncSubHandlerNotFound(t *testing.T) {
	var testServer = NewServer()
	err := testServer.PutFunc("/mufasa", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		t.Log("PUT")
	})

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

func TestServerPutFuncMethodNotFound(t *testing.T) {
	var testServer = NewServer()
	err := testServer.PutFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		t.Log("PUT")
	})

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

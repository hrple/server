package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServerStart(t *testing.T) {
	go Start(":5868")
}

func TestServerGetConfiguration(t *testing.T) {
	config := GetConfiguration()

	if config.ReadTimeoutInSeconds != DefaultWriteTimeOutInSeconds {
		t.Fatalf("Something wrong")
	}
}

func TestServerGetLogger(t *testing.T) {
	logger := GetLogger()
	if logger == nil {
		t.Fatalf("Something wrong")
	}
}

func TestServerGet(t *testing.T) {
	Get("/testGet", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	request, _ := http.NewRequest(http.MethodGet, "/testGet", nil)
	response := httptest.NewRecorder()
	applicationServer.router.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("Something wrong")
	}
}

func TestServerPut(t *testing.T) {
	Get("/testPut", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	request, _ := http.NewRequest(http.MethodPut, "/testPut", nil)
	response := httptest.NewRecorder()

	applicationServer.router.ServeHTTP(response, request)

	if response.Code != http.StatusNotFound {
		t.Fatalf("Something wrong")
	}
}


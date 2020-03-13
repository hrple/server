package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServerStart(t *testing.T) {
	serviceRunning := make(chan struct{})
	serviceDone := make(chan struct{})

	Get("/testGet", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	var err error
	var errStop error
	go func() {
		close(serviceRunning)
		err = Start(":5868")

		defer close(serviceDone)
	}()

	<-serviceRunning

	if err != nil {
		t.Fatalf("Server never started %v", err)
	}
	errStop = Stop()
	if errStop != nil {
		t.Fatalf("Server never started %v", errStop)
	}

	<-serviceDone
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

var testServer = NewServer()

func TestServerGet(t *testing.T) {
	testServer.Get("/test", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		t.Log("GET")
	})

	request, _ := http.NewRequest(http.MethodGet, "/test", nil)
	response := httptest.NewRecorder()

	testServer.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("Something wrong, code : %v", response.Code)
	}
}

func TestServerPut(t *testing.T) {
	testServer.Put("/test", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		t.Log("PUT")
	})

	request, _ := http.NewRequest(http.MethodPut, "/test", nil)
	response := httptest.NewRecorder()

	testServer.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("Something wrong")
	}
}

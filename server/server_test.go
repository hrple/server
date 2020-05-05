package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServerStart(t *testing.T) {
	serviceRunning := make(chan struct{})
	serviceDone := make(chan struct{})

	err := Get("/testGet", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	if err != nil {
		t.Errorf("Error Occurred: %v", err)
	}

	err = Put("/testGet", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	if err != nil {
		t.Errorf("Error Occurred: %v", err)
	}

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

func TestServerGet(t *testing.T) {
	var testServer = NewServer()
	err := testServer.Get("/test", func(w http.ResponseWriter, r *http.Request) {
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

func TestServerPut(t *testing.T) {
	var testServer = NewServer()
	err := testServer.Put("/test", func(w http.ResponseWriter, r *http.Request) {
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

func TestServerPutHandlerNotFound(t *testing.T) {
	var testServer = NewServer()
	err := testServer.Put("/test", func(w http.ResponseWriter, r *http.Request) {
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

func TestServerPutSubHandlerNotFound(t *testing.T) {
	var testServer = NewServer()
	err := testServer.Put("/mufasa", func(w http.ResponseWriter, r *http.Request) {
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

func TestServerPutMethodNotFound(t *testing.T) {
	var testServer = NewServer()
	err := testServer.Put("/test", func(w http.ResponseWriter, r *http.Request) {
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

func TestServerAddInvalidRoute(t *testing.T) {
	var testServer = NewServer()
	err := testServer.Get("[ INVALID REGEX", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	if err == nil {
		t.Error("This should have failed an an invalid regex was passed int to GET")
	}
}

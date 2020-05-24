package server

import (
	"net/http"
	"testing"
)

func TestServerStart(t *testing.T) {
	serviceRunning := make(chan struct{})
	serviceDone := make(chan struct{})

	err := GetFunc("/testGet", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	if err != nil {
		t.Errorf("Error Occurred: %v", err)
	}

	err = PutFunc("/testPut", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	if err != nil {
		t.Errorf("Error Occurred: %v", err)
	}

	err = Put("/testPut", &HrpleTestHandler{})

	if err != nil {
		t.Errorf("Error Occurred: %v", err)
	}

	err = Get("/testGet", &HrpleTestHandler{})

	if err != nil {
		t.Errorf("Error Occurred: %v", err)
	}

	var errStop error
	go func() {
		close(serviceRunning)
		err = Start("")

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

func TestServerStartIPv4(t *testing.T) {
	serviceRunning := make(chan struct{})
	serviceDone := make(chan struct{})

	err := GetFunc("/testGet", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	if err != nil {
		t.Errorf("Error Occurred: %v", err)
	}

	err = PutFunc("/testPut", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	if err != nil {
		t.Errorf("Error Occurred: %v", err)
	}

	var errStop error
	go func() {
		close(serviceRunning)
		err = Start("127.0.0.1:8080")

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

func TestServerFuncAddInvalidRoute(t *testing.T) {
	var testServer = NewServer()
	err := testServer.GetFunc("[ INVALID REGEX", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	if err == nil {
		t.Error("This should have failed an an invalid regex was passed int to GET")
	}
}

func TestServerAddInvalidRoute(t *testing.T) {
	var testServer = NewServer()
	err := testServer.Get("[ INVALID REGEX", &HrpleTestHandler{})

	if err == nil {
		t.Error("This should have failed an an invalid regex was passed int to GET")
	}
}

type HrpleTestHandler struct {
}

func (h *HrpleTestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

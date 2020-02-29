package server

import (
	"log"
	"os"
	"strings"
	"testing"
	"time"
)

var (
	SampleAPIServiceAddress = os.Getenv("SAMPLE_API_SERVICE_ADDR")
	SampleAPIServerCert     = os.Getenv("SAMPLE_API_TLS_CERT_FILE")
	SampleAPIServerKey      = os.Getenv("SAMPLE_API_TLS_KEY_FILE")
	SampleAPIRunningContext = os.Getenv("SAMPLE_API_RUNNING_CONTEXT")
)

const ApplicationName = "HRPLE-SERVER-TEST"

func TestStandaloneServerInitialisation(t *testing.T) {
	defaultReadTimeout, _ := time.ParseDuration("5s")
	defaultWriteTimeout, _ := time.ParseDuration("10s")
	defaultIdleTimeout, _ := time.ParseDuration("120s")

	var runningContextType, err = GetRunningContextType(SampleAPIRunningContext)
	if err != nil {
		if strings.Contains(err.Error(), "WARNING:") {
			t.Logf("%v", err)
		} else {
			t.Fatalf("Error failed to GetRunningContextType - Error: %v", err)
		}
	}

	appServerConfig := &ApplicationServerConfig{
		ServerAddress: SampleAPIServiceAddress,
		TLSCertFile:   SampleAPIServerCert,
		TLSKeyFile:    SampleAPIServerKey,
		ReadTimeout:   defaultReadTimeout,
		WriteTimeout:  defaultWriteTimeout,
		IdleTimeout:   defaultIdleTimeout,
	}

	logPrefix := ApplicationName + " : "
	logger := log.New(os.Stdout, logPrefix, log.LstdFlags|log.Lshortfile)

	appServer, err := New(runningContextType, logger, appServerConfig)
	if err != nil {
		t.Fatal("Error failed to init server")
	}

	t.Log(appServer)

	//appServer.Run()

	// cases := []struct {
	// 	name        string
	// 	path        string
	// 	contentType string
	// }{
	// 	{
	// 		name:        "normal file",
	// 		path:        "index.html",
	// 		contentType: "",
	// 	},
	// 	{
	// 		name:        "javascript",
	// 		path:        "test.js",
	// 		contentType: "application/javascript",
	// 	},
	// 	{
	// 		name:        "css",
	// 		path:        "test.css",
	// 		contentType: "text/css",
	// 	},
	// 	{
	// 		name:        "png",
	// 		path:        "test.png",
	// 		contentType: "image/png",
	// 	},
	// 	{
	// 		name:        "jpg",
	// 		path:        "test.jpg",
	// 		contentType: "image/jpeg",
	// 	},
	// 	{
	// 		name:        "gif",
	// 		path:        "test.gif",
	// 		contentType: "image/gif",
	// 	},
	// }

	// for _, c := range cases {
	// 	t.Run(c.name, func(t *testing.T) {
	// 		rr := httptest.NewRecorder()
	// 		req, err := http.NewRequest("GET", "http://localhost/"+c.path, nil)

	// 		if err != nil {
	// 			t.Fatal(err)
	// 		}

	// 		s := StaticFileServer(dummyFileSystem{})
	// 		s.ServeHTTP(rr, req)

	// 		if rr.Header().Get("Content-Type") != c.contentType {
	// 			t.Fatalf("Unexpected Content-Type: %s", rr.Header().Get("Content-Type"))
	// 		}
	// 	})
	// }
}

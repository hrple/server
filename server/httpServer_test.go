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
)

const ApplicationName = "HRPLE-SERVER-TEST"

func getAppServerConfig() *ApplicationServerConfig {
	defaultReadTimeout, _ := time.ParseDuration("5s")
	defaultWriteTimeout, _ := time.ParseDuration("10s")
	defaultIdleTimeout, _ := time.ParseDuration("120s")

	appServerConfig := &ApplicationServerConfig{
		ServerAddress: SampleAPIServiceAddress,
		TLSCertFile:   SampleAPIServerCert,
		TLSKeyFile:    SampleAPIServerKey,
		ReadTimeout:   defaultReadTimeout,
		WriteTimeout:  defaultWriteTimeout,
		IdleTimeout:   defaultIdleTimeout,
	}

	return appServerConfig
}

func getLogger() *log.Logger {
	logPrefix := ApplicationName + " : "
	logger := log.New(os.Stdout, logPrefix, log.LstdFlags|log.Lshortfile)
	return logger
}

func TestStandaloneServerInitialisationLambda(t *testing.T) {

	var runningContextType, err = GetRunningContextType(RunningContextTypeLambda)
	if err != nil {
		if strings.Contains(err.Error(), "WARNING:") {
			t.Logf("%v", err)
		} else {
			t.Fatalf("Error failed to GetRunningContextType - Error: %v", err)
		}
	}

	appServerConfig := getAppServerConfig()
	logger := getLogger()

	appServer, err := New(runningContextType, logger, appServerConfig)
	if err != nil {
		//t.Fatal("Error failed to init server")
		t.Logf("Not implemented - Error : %v", err)
	}

	if appServer == nil {
		t.Fatal("Error failed to init server, expected AppServer object")
	}
}

func TestStandaloneServerInitialisationStandalone(t *testing.T) {
	var runningContextType, err = GetRunningContextType(RunningContextTypeStandalone)
	if err != nil {
		if strings.Contains(err.Error(), "WARNING:") {
			t.Logf("%v", err)
		} else {
			t.Fatalf("Error failed to GetRunningContextType - Error: %v", err)
		}
	}

	appServerConfig := getAppServerConfig()
	logger := getLogger()

	appServer, err := New(runningContextType, logger, appServerConfig)
	if err != nil {
		t.Fatal("Error failed to init server")
	}

	if appServer == nil {
		t.Fatal("Error failed to init server, expected AppServer object")
	}
}

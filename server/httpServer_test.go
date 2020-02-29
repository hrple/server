package server

import (
	"context"
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

const (
	ApplicationName = "HRPLE-SERVER-TEST"
	Warning         = "WARNING"
)

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

func TestServerInitialisationLambda(t *testing.T) {
	var runningContextType, err = GetRunningContextType(RunningContextTypeLambda)
	if err != nil {
		if strings.Contains(err.Error(), Warning) {
			t.Logf("%v", err)
		} else {
			t.Fatalf(ErrMsgExpectedInsteadOfResultWithError, RunningContextTypeLambda, runningContextType, err)
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

func TestServerInitialisationStandalone(t *testing.T) {
	var runningContextType, err = GetRunningContextType(RunningContextTypeStandalone)
	if err != nil {
		if strings.Contains(err.Error(), Warning) {
			t.Logf("%v", err)
		} else {
			t.Fatalf(ErrMsgExpectedInsteadOfResultWithError, RunningContextTypeStandalone, runningContextType, err)
		}
	}

	appServerConfig := getAppServerConfig()
	logger := getLogger()

	appServer, err := New(runningContextType, logger, appServerConfig)
	if err != nil {
		t.Fatalf(ErrMsgFailedCreateServerWithError, err)
	}

	if appServer == nil {
		t.Fatal(ErrMsgFailedCreateServer)
	}
}

func TestRunStandaloneServer(t *testing.T) {
	var runningContextType, err = GetRunningContextType(RunningContextTypeStandalone)
	if err != nil {
		if strings.Contains(err.Error(), Warning) {
			t.Logf("%v", err)
		} else {
			t.Fatalf(ErrMsgExpectedInsteadOfResultWithError, RunningContextTypeStandalone, runningContextType, err)
		}
	}

	appServerConfig := getAppServerConfig()
	logger := getLogger()

	appServer, err := New(runningContextType, logger, appServerConfig)
	if err != nil {
		t.Fatalf(ErrMsgFailedCreateServerWithError, err)
	}

	if appServer == nil {
		t.Fatal(ErrMsgFailedCreateServer)
	}

	serviceRunning := make(chan struct{})
	serviceDone := make(chan struct{})
	go func() {
		close(serviceRunning)
		err = appServer.Run()

		defer close(serviceDone)
	}()

	if err != nil {
		t.Fatal("Server never started")
	}

	err = appServer.httpServer.Shutdown(context.Background())
	if err != nil {
		t.Fatal("Server never shutdown")
	}
}

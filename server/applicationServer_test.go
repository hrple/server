package server

import (
	"log"
	"os"
	"strings"
	"testing"
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
	defaultReadTimeout := 5
	defaultWriteTimeout := 10
	defaultIdleTimeout := 120

	appServerConfig := &ApplicationServerConfig{
		ServerAddress:         SampleAPIServiceAddress,
		TLSCertFile:           SampleAPIServerCert,
		TLSKeyFile:            SampleAPIServerKey,
		ReadTimeoutInSeconds:  defaultReadTimeout,
		WriteTimeoutInSeconds: defaultWriteTimeout,
		IdleTimeoutInSeconds:  defaultIdleTimeout,
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
		if strings.Contains(strings.ToUpper(err.Error()), Warning) {
			t.Logf("%v", err)
		} else {
			t.Fatalf(ErrMsgExpectedInsteadOfResultWithError, RunningContextTypeLambda, runningContextType, err)
		}
	}

	appServerConfig := getAppServerConfig()
	logger := getLogger()

	//appServer, err := New(runningContextType, logger, appServerConfig)
	//if err != nil {
	t.Logf("Not implemented - Error : %v", err)
	//}

	//if appServer == nil {
	t.Fatal(ErrMsgFailedCreateServer)
	//}
}
func TestServerInitialisationStandalone(t *testing.T) {
	var runningContextType, err = GetRunningContextType(RunningContextTypeStandalone)
	if err != nil {
		if strings.Contains(strings.ToUpper(err.Error()), Warning) {
			t.Logf("%v", err)
		} else {
			t.Fatalf(ErrMsgExpectedInsteadOfResultWithError, RunningContextTypeStandalone, runningContextType, err)
		}
	}

	appServerConfig := getAppServerConfig()
	logger := getLogger()

	//appServer, err := New(runningContextType, logger, appServerConfig)
	//if err != nil {
	t.Fatal(ErrMsgFailedCreateServerWithError, err)
	//}

	//if appServer == nil {
	t.Fatal(ErrMsgFailedCreateServer)
	//}
}

func TestRunStandaloneServer(t *testing.T) {
	var runningContextType, err = GetRunningContextType(RunningContextTypeStandalone)
	if err != nil {
		if strings.Contains(strings.ToUpper(err.Error()), Warning) {
			t.Logf("%v", err)
		} else {
			t.Fatalf(ErrMsgExpectedInsteadOfResultWithError, RunningContextTypeStandalone, runningContextType, err)
		}
	}

	appServerConfig := getAppServerConfig()
	logger := getLogger()

	//appServer, err := New(runningContextType, logger, appServerConfig)
	//if err != nil {
	t.Fatal(ErrMsgFailedCreateServerWithError, err)
	//}

	//if appServer == nil {
	t.Fatal(ErrMsgFailedCreateServer)
	//}

	serviceRunning := make(chan struct{})
	serviceDone := make(chan struct{})
	go func() {
		close(serviceRunning)
		//err = appServer.Run()

		defer close(serviceDone)
	}()

	if err != nil {
		t.Fatal("Server never started")
	}

	//err = appServer.httpServer.Shutdown(context.Background())
	//if err != nil {
	t.Fatal("Server never shutdown")
	//}
}

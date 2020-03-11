package server

import (
	"log"
	"net/http"
	"os"
)

// Default values for configuration
const (
	DefaultWriteTimeOutInSeconds = 10
	DefaultIdleTimeOutInSeconds  = 120
	DefaultServerAddress         = ""
)

// ApplicationServerConfig allows for the configuration of the App Server
type ApplicationServerConfig struct {
	ServerAddress         string
	TLSCertFile           string
	TLSKeyFile            string
	ReadTimeoutInSeconds  int
	WriteTimeoutInSeconds int
	IdleTimeoutInSeconds  int
}

// ApplicationServer is a wrapper for the required loggers, handlers, https server etc
type ApplicationServer struct {
	logger     *log.Logger
	mux        *http.ServeMux
	httpServer *http.Server
	config     *ApplicationServerConfig
}

// Config the default configuration for the server
var Config = &ApplicationServerConfig{
	ServerAddress:         "",
	TLSCertFile:           "",
	TLSKeyFile:            "",
	ReadTimeoutInSeconds:  DefaultWriteTimeOutInSeconds,
	WriteTimeoutInSeconds: DefaultWriteTimeOutInSeconds,
	IdleTimeoutInSeconds:  DefaultIdleTimeOutInSeconds,
}

// NewServer creates an instance of the app server
func NewServer() *ApplicationServer {
	return &ApplicationServer{
		config: Config,
		logger: log.New(os.Stdout, "", log.Ldate|log.Ltime),
	}
}

// Start will start the server eventually
func (s *ApplicationServer) Start() error {
	return nil
}

var applicationServer = NewServer()

// New instantiates a new ApplicationServer based on the ApplicationServerConfig
// func New(runningContextType string, logger *log.Logger, config *ApplicationServerConfig) (*ApplicationServer, error) {
// 	var err error

// 	server := &ApplicationServer{
// 		logger:     logger,
// 		mux:        nil,
// 		httpServer: nil,
// 		config:     config,
// 	}

// 	if runningContextType == RunningContextTypeLambda {
// 		err = errors.New("warning lambda not operational yet")
// 	}

// 	if runningContextType == RunningContextTypeStandalone {
// 		// TODO err = errors.New("Warning. TLS not operational yet")
// 		mux := http.NewServeMux()

// 		httpServer := &http.Server{
// 			Addr:         config.ServerAddress,
// 			ReadTimeout:  config.ReadTimeout,
// 			WriteTimeout: config.WriteTimeout,
// 			IdleTimeout:  config.IdleTimeout,
// 			// TLSConfig:    tlsConfig,
// 			Handler: mux,
// 		}

// 		server.mux = mux
// 		server.httpServer = httpServer
// 	}

// 	return server, err
// }

// // Run starts the ApplicationServer
// func (server *ApplicationServer) Run() error {
// 	server.logger.Println("Info -", "Run")
// 	err := server.httpServer.ListenAndServe()

// 	return err
// }

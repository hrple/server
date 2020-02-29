package server

import (
	"errors"
	"log"
	"net/http"
	"time"
)

//ApplicationServerConfig allows for the configuration of the App Server
type ApplicationServerConfig struct {
	ServerAddress string
	TLSCertFile   string
	TLSKeyFile    string
	ReadTimeout   time.Duration
	WriteTimeout  time.Duration
	IdleTimeout   time.Duration
}

//ApplicationServer is a wrapper for the required loggers, handlers, https server etc
type ApplicationServer struct {
	logger     *log.Logger
	mux        *http.ServeMux
	httpServer *http.Server
	config     *ApplicationServerConfig
}

//Run starts the ApplicationServer
func (server *ApplicationServer) Run() error {
	server.logger.Println("Info -", "Run")
	err := server.httpServer.ListenAndServe()

	return err
}

//New instantiates a new ApplicationServer based on the ApplicationServerConfig
func New(runningContextType string, logger *log.Logger, config *ApplicationServerConfig) (*ApplicationServer, error) {
	var err error

	server := &ApplicationServer{
		logger:     logger,
		mux:        nil,
		httpServer: nil,
		config:     config,
	}

	if runningContextType == RunningContextTypeLambda {
		err = errors.New("Warning. LAMBDA not operational yet")
	}

	if runningContextType == RunningContextTypeStandalone {
		//err = errors.New("Warning. TLS not operational yet")
		mux := http.NewServeMux()

		httpServer := &http.Server{
			Addr:         config.ServerAddress,
			ReadTimeout:  config.ReadTimeout,
			WriteTimeout: config.WriteTimeout,
			IdleTimeout:  config.IdleTimeout,
			// TLSConfig:    tlsConfig,
			Handler: mux,
		}

		server.mux = mux
		server.httpServer = httpServer
	}

	return server, err
}

// func New(mux *http.ServeMux, serverAddress string) *http.Server {
// 	tlsConfig := &tls.Config{
// 		// Causes servers to use Go's default ciphersuite preferences,
// 		// which are tuned to avoid attacks. Does nothing on clients.
// 		PreferServerCipherSuites: true,
// 		// Only use curves which have assembly implementations
// 		CurvePreferences: []tls.CurveID{
// 			tls.CurveP256,
// 			tls.X25519, // Go 1.8 only
// 		},
// 		MinVersion: tls.VersionTLS12,
// 		CipherSuites: []uint16{
// 			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
// 			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
// 			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305, // Go 1.8 only
// 			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,   // Go 1.8 only
// 			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
// 			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
// 		},
// 	}

// 	srv := &http.Server{
// 		Addr:         serverAddress,
// 		ReadTimeout:  5 * time.Second,
// 		WriteTimeout: 10 * time.Second,
// 		IdleTimeout:  120 * time.Second,
// 		TLSConfig:    tlsConfig,
// 		Handler:      mux,
// 	}

// 	return srv
// }

package server

import (
	"log"
	"net/http"
)

// Start will start the server eventually
func Start(address string) error {
	err := applicationServer.Start(address)
	return err
}

// GetConfiguration gets the cofiguration
func GetConfiguration() *ApplicationServerConfig {
	return applicationServer.Config
}

// GetLogger gets the logger
func GetLogger() *log.Logger {
	return applicationServer.Logger
}

// Get adds a handler for the 'GET' http method for server s.
func Get(route string, f func(http.ResponseWriter, *http.Request)) {
	applicationServer.router.HandleFunc(route, applicationServer.get(f))
}

var applicationServer = NewServer()

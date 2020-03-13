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

// Stop will stop the server eventually
func Stop() error {
	err := applicationServer.Stop()
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
	applicationServer.Get(route, f)
}

var applicationServer = NewServer()

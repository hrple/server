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
func Get(route string, h http.Handler) error {
	return applicationServer.Get(route, h)
}

// Put adds a handler for the 'PUT' http method for server s.
func Put(route string, h http.Handler) error {
	return applicationServer.Put(route, h)
}

// GetFunc adds a handler for the 'GET' http method for server s.
func GetFunc(route string, f func(http.ResponseWriter, *http.Request)) error {
	return applicationServer.GetFunc(route, f)
}

// PutFunc adds a handler for the 'PUT' http method for server s.
func PutFunc(route string, f func(http.ResponseWriter, *http.Request)) error {
	return applicationServer.GetFunc(route, f)
}

var applicationServer = NewServer()

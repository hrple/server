package server

import (
	"net/http"
	"testing"
)

func TestApplicationServerGet(t *testing.T) {
	f := applicationServer.get(func(w http.ResponseWriter, r *http.Request) {})
	if f == nil {
		t.Fatalf("Something wrong")
	}
}

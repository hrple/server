package main

import (
	"fmt"
	"github.com/hrple/common/server"
	"net/http"
)

func main() {
	server.Get("/hello/", helloWorldStd)
	logger := server.GetLogger()
	logger.Fatal(server.Start(""))
}

func helloWorldStd(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Hello World - Simple Server %v \n", r.URL)
}

package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"
)

// NewServer creates an instance of the app server
func NewServer() *ApplicationServer {
	server := &ApplicationServer{
		Config: Config,
		Logger: log.New(os.Stdout, "", log.Ldate|log.Ltime),
		router: http.NewServeMux(),
	}

	server.httpServer = http.Server{
		Addr:    server.Config.ServerAddress,
		Handler: server.router,
	}

	return server
}

// Start will start the server eventually
func (s *ApplicationServer) Start(addr string) error {
	s.Config.ServerAddress = addr
	s.httpServer.Addr = s.Config.ServerAddress

	err := s.httpServer.ListenAndServe()
	return err
}

// Stop will stop the server eventually
func (s *ApplicationServer) Stop() error {
	err := s.httpServer.Shutdown(context.Background())
	return err
}

// Get adds a handler for the 'GET' http method for server s.
func (s *ApplicationServer) Get(route string, f func(http.ResponseWriter, *http.Request)) {
	s.router.HandleFunc(route, s.get(f))
}

func (s *ApplicationServer) get(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctime := time.Now()

		defer s.Logger.Printf("Executed: %v%v Execution Time: %v Method : %v RemoteAddr: %v",
			r.Host, r.URL, time.Since(ctime), r.Method, r.RemoteAddr)

		// should maybe add this later s.before(w, r)

		if r.Method == http.MethodGet {
			h(w, r)
		} else {
			http.NotFound(w, r)
		}

		// should maybe add this later  s.after(w, r)
	}
}

//	func (s *ApplicationServer) process(h http.HandlerFunc) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		ctime := time.Now()
//		defer s.Logger.Printf("Executed: %v%v Execution Time: %v Method : %v RemoteAddr: %v",
//			r.Host, r.URL, time.Since(ctime), r.Method, r.RemoteAddr)
//		s.before(w, r)
//		h(w, r)
//		s.after(w, r)
//	}
//	}
//
// func (s *ApplicationServer) before(w http.ResponseWriter, r *http.Request) {
//	// _, _ = fmt.Fprintf(w, "before %v \n", r.URL)
//	s.Logger.Printf("in before method \n")
// }

// url := r.URL.String()
// path := strings.Split(url, "/")

// Add this for basic param injection {userId}, add route will need to be implemented, but for now this is fine.
// for _, element := range path {
//	if (element != "std" && element != "") {
//		fmt.Fprintf(w, "My path Element is: %v \n", element)
//	}
// }
// func (s *ApplicationServer) after(w http.ResponseWriter, r *http.Request) {
//	// _, _ = fmt.Fprintf(w, "after %v \n", r.URL)
//	s.Logger.Printf("in after method \n")
// }

// TODO Add middle ware for logging and time taken to process, then add stats ie avg time taken
// ServeHTTP is the interface method for Go's http server package
// func (s *ApplicationServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	 s.Process(c, req)
//	w.Write([]byte("foo ServeHTTP"))
// }

// Default values for configuration
const (
	DefaultWriteTimeOutInSeconds = 10
	DefaultIdleTimeOutInSeconds  = 120
	// DefaultServerAddress         = ""
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
	Logger     *log.Logger
	router     *http.ServeMux
	Config     *ApplicationServerConfig
	httpServer http.Server
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

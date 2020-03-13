package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"regexp"
	"time"
)

// NewServer creates an instance of the app server
func NewServer() *ApplicationServer {
	server := &ApplicationServer{
		Config: Config,
		Logger: log.New(os.Stdout, "", log.Ldate|log.Ltime),
		router: http.NewServeMux(),
	}

	server.router.Handle("/", server)

	server.httpServer = http.Server{
		Addr:    server.Config.ServerAddress,
		Handler: server.router,
	}

	// MaxIdleConns:       10,
	// IdleConnTimeout:    30 * time.Second,
	// DisableCompression: true,

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
	s.addRoute(route, http.MethodGet, f)
}

// Put adds a handler for the 'PUT' http method for server s.
func (s *ApplicationServer) Put(route string, f func(http.ResponseWriter, *http.Request)) {
	s.addRoute(route, http.MethodPut, f)
}

// ServeHTTP is the interface method for Go's http server package
func (s *ApplicationServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctime := time.Now()
	route := s.findRouteHandler(w, r)
	if route == nil {
		http.NotFound(w, r)
		return
	}

	w.Header().Add("Content-Type", "text/plain; charset=utf-8")

	// should maybe add this later s.before(w, r)
	// should maybe add this later  s.after(w, r)
	route.httpHandler(w, r)
	s.Logger.Printf("Executed: %v%v Execution Time: %v Method : %v RemoteAddr: %v",
		r.Host, r.URL, time.Since(ctime), r.Method, r.RemoteAddr)
}

func (s *ApplicationServer) findRouteHandler(w http.ResponseWriter, r *http.Request) (selectedRoute *route) {
	requestPath := r.URL.Path

	if w == nil {
		return nil
	}

	for i := 0; i < len(s.routes); i++ {
		currentRoute := s.routes[i]
		cr := currentRoute.cr

		// if the methods don't match, skip this handler
		if r.Method != currentRoute.method {
			continue
		}

		if !cr.MatchString(requestPath) {
			continue
		}

		match := cr.FindStringSubmatch(requestPath)

		if len(match[0]) != len(requestPath) {
			continue
		}

		if currentRoute.httpHandler != nil {
			selectedRoute = &currentRoute
			return selectedRoute
		}

		return nil
	}

	return nil
}

func (s *ApplicationServer) addRoute(r, method string, handler func(http.ResponseWriter, *http.Request)) {
	cr, err := regexp.Compile(r)
	if err != nil {
		s.Logger.Printf("Error in route regex %q\n", r)
		return
	}

	s.routes = append(s.routes, route{r: r, cr: cr, method: method, httpHandler: handler})
}

// func (s *ApplicationServer) before(w http.ResponseWriter, r *http.Request) {
//	// _, _ = fmt.Fprintf(w, "before %v \n", r.URL)
//	s.Logger.Printf("in before method \n")
// }

// func (s *ApplicationServer) after(w http.ResponseWriter, r *http.Request) {
//	// _, _ = fmt.Fprintf(w, "after %v \n", r.URL)
//	s.Logger.Printf("in after method \n")
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
	routes     []route
	Config     *ApplicationServerConfig
	httpServer http.Server
}

type route struct {
	r           string
	cr          *regexp.Regexp
	method      string
	httpHandler func(http.ResponseWriter, *http.Request)
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

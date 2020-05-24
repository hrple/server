package server

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"os"
	"regexp"
	"strings"
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

	listner, err := net.Listen("tcp", s.Config.ServerAddress)
	if err != nil {
		log.Fatal("Listen:", err)
	} else {
		address := listner.Addr().String()
		ip := ""
		port := ""

		arr := strings.Split(address, ":")
		if strings.Contains(address, "[::]") {
			ip = "127.0.0.1"
			port = arr[len(arr)-1]
		} else {
			// IPV4
			const numberOfItems int = 2
			if len(arr) == numberOfItems {
				ip = arr[0]
				port = arr[1]
			}

			if len(arr) > numberOfItems {
				// IPV6
				log.Fatal("IPV6 Not Implemented, bind to IPV4")
			}
		}

		log.Println("Listening on:", address, " URL: http://"+ip+":"+port)

		if s.Config.ServerAddress == "" {
			s.Config.ServerAddress = address
		}
	}

	err = s.httpServer.Serve(listner)
	return err
}

// Stop will stop the server eventually
func (s *ApplicationServer) Stop() error {
	err := s.httpServer.Shutdown(context.Background())
	return err
}

// Get adds a handler for the 'GET' http method for server s.
func (s *ApplicationServer) Get(route string, h http.Handler) error {
	err := s.addRoute(route, http.MethodGet, h)
	return err
}

// GetFunc adds a handler for the 'GET' http method for server s.
func (s *ApplicationServer) GetFunc(route string, f func(http.ResponseWriter, *http.Request)) error {
	err := s.addRouteFunc(route, http.MethodGet, f)
	return err
}

// Put adds a handler for the 'PUT' http method for server s.
func (s *ApplicationServer) Put(route string, h http.Handler) error {
	err := s.addRoute(route, http.MethodPut, h)
	return err
}

// PutFunc adds a handler for the 'PUT' http method for server s.
func (s *ApplicationServer) PutFunc(route string, f func(http.ResponseWriter, *http.Request)) error {
	err := s.addRouteFunc(route, http.MethodPut, f)
	return err
}

// ServeHTTP is the interface method for Go's http server package
func (s *ApplicationServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctime := time.Now()
	// remove w, will add it if its needed.
	route := s.findRouteHandler(r)

	if route == nil {
		http.NotFound(w, r)
		s.Logger.Printf("Info: Route Not Found (%v)", r.URL)
		return
	}

	// add common headers
	w.Header().Add("Server", "hrple.common.server")
	w.Header().Add("Content-Type", "text/html; charset=utf-8")

	// remove before and after, will add it if its needed s.before(w, r) / s.after(w, r)
	if route.httpHandler != nil {
		route.httpHandler.ServeHTTP(w, r)
	} else {
		route.httpHandlerFunc(w, r)
	}

	s.Logger.Printf("Executed: %v%v Execution Time: %v Method : %v RemoteAddr: %v",
		r.Host, r.URL, time.Since(ctime), r.Method, r.RemoteAddr)
}

func (s *ApplicationServer) findRouteHandler(r *http.Request) (selectedRoute *route) {
	requestPath := r.URL.Path

	// remove trailing slash if any (i.e. GET /hello/ equals GET /hello)
	// lastChar := requestPath[len(requestPath)-1:]
	// if lastChar == "/" {
	// 	requestPath = requestPath[:len(requestPath)-1]
	// }

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

		if currentRoute.httpHandler != nil || currentRoute.httpHandlerFunc != nil {
			selectedRoute = &currentRoute
			break
		}
	}

	return selectedRoute
}

// func (s *ApplicationServer) addRoute(r, method string, handler func(http.ResponseWriter, *http.Request)) error {
// 	cr, err := regexp.Compile(r)
// 	if err != nil {
// 		s.Logger.Printf("Error in route regex %q\n", r)
// 		err = errors.New("invalid reg expression, unable to add route")
// 		return err
// 	}

// 	s.routes = append(s.routes, route{r: r, cr: cr, method: method, httpHandler: handler})
// 	return nil
// }

func (s *ApplicationServer) addRoute(r, method string, h http.Handler) error {
	cr, err := s.checkRoute(r)
	if err != nil {
		return err
	}

	s.routes = append(s.routes, route{r: r, cr: cr, method: method, httpHandler: h})
	return nil
}

func (s *ApplicationServer) addRouteFunc(r, method string, f func(w http.ResponseWriter, r *http.Request)) error {
	cr, err := s.checkRoute(r)
	if err != nil {
		return err
	}

	s.routes = append(s.routes, route{r: r, cr: cr, method: method, httpHandlerFunc: f})
	return nil
}

func (s *ApplicationServer) checkRoute(r string) (*regexp.Regexp, error) {
	cr, err := regexp.Compile(r)
	if err != nil {
		s.Logger.Printf("Error in route regex %q\n", r)
		err = errors.New("invalid reg expression, unable to add route")
		return nil, err
	}

	return cr, err
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
	r               string
	cr              *regexp.Regexp
	method          string
	httpHandler     http.Handler
	httpHandlerFunc func(http.ResponseWriter, *http.Request)
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

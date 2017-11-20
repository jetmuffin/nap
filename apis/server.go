package api

import (
	"net"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/JetMuffin/nap/pkg/api/router"
	"github.com/Sirupsen/logrus"
	"github.com/JetMuffin/nap/pkg/api/router/debug"
)

// Config provides the configuration for the API server
type Config struct {
	LogLevel    string
	CorsHeaders string
	Version     string
}

// Server contains instance details for the server
type Server struct {
	cfg           *Config
	server        *http.Server
	listener      net.Listener
	routers       []router.Router
	routerWrapper *mux.Router
}

func New(cfg *Config) *Server {
	s := &Server{
		cfg:      cfg,
	}
	return s
}

func (s *Server) Accept(addr string, listener net.Listener) {
	s.server = &http.Server{
		Addr: addr,
	}
	s.listener = listener
}

func (s *Server) InitRouter(routers ...router.Router) {
	s.routers = append(s.routers, routers...)

	s.routerWrapper = s.createMux()
}

// createMux initializes the main router the server uses.
func (s *Server) createMux() *mux.Router {
	m := mux.NewRouter()

	logrus.Debug("Registering routers")
	for _, apiRouter := range s.routers {
		for _, r := range apiRouter.Routes() {
			handler := s.makeHTTPHandler(r.Handler())

			m.Path(r.Path()).Methods(r.Method()).Handler(handler)
			logrus.Debugf("Registering %s, %s", r.Method(), r.Path())
		}
	}

	debugRouter := debug.NewRouter()
	s.routers = append(s.routers, debugRouter)
	for _, r := range debugRouter.Routes() {
		handler := s.makeHTTPHandler(r.Handler())
		m.Path("/debug" + r.Path()).Handler(handler)
	}

	return m
}

func (s *Server) enableCORS(w http.ResponseWriter) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, X-Registry-Auth")
	w.Header().Add("Access-Control-Allow-Methods", "HEAD, GET, POST, DELETE, PUT, OPTIONS")
}

func (s *Server) makeHTTPHandler(handler router.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.enableCORS(w)

		handler(w, r)
	}
}

func (s *Server) serve() error {
	s.server.Handler = s.routerWrapper
	var chErrors = make(chan error, 1)
	go func() {
		var err error
		logrus.Infof("API server listen on %s", s.listener.Addr())
		if err = s.server.Serve(s.listener); err != nil {
			chErrors <- err
		}
	}()
	err := <-chErrors
	if err != nil {
		return err
	}
	return nil
}

// gracefully shutdown.
func (s *Server) Shutdown() error {
	// If s.server is nil, api server is not running.
	if s.server != nil {
		// NOTE(nmg): need golang 1.8+ to run this method.
		return s.server.Shutdown(nil)
	}

	return nil
}

func (s *Server) Stop() error {
	if s.server != nil {
		return s.server.Close()
	}

	return nil
}

// Wait blocks the server goroutine until it exits.
// It sends an error message if there is any error during
// the API execution.
func (s *Server) Wait(waitChan chan error) {
	if err := s.serve(); err != nil {
		logrus.Errorf("ServeAPI error: %v", err)
		waitChan <- err
		return
	}
	waitChan <- nil
}
package app

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"sprinteins.com/go-bank/src/x/log"

	"github.com/gorilla/mux"
)

//
// Server provides an http server wrap around services
type Server struct {
	host    string
	port    string
	stop    chan os.Signal
	stopped chan struct{}
	server  *http.Server
}

// NewServer _
func NewServer(
	host string,
	port string,
) (s *Server) {

	return &Server{
		host:    host,
		port:    port,
		stop:    make(chan os.Signal, 1),
		stopped: make(chan struct{}, 1),
	}
}

// Start _
func (s *Server) Start() func() chan struct{} {

	go func() {
		var router = mux.NewRouter()

		router.HandleFunc("/", rootHandler)

		var address = fmt.Sprintf("%s:%s", s.host, s.port)
		s.server = &http.Server{Addr: address, Handler: router}

		log.Info.Printf("state=http_listening address=http://%s", address)
		go func() {
			if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Error.Println(err)
			}
		}()
		signal.Notify(s.stop, os.Interrupt, syscall.SIGTERM)
		s.waitForStop(&s.stop, s.server)
	}()

	return s.Stop
}

// Address _
func (s *Server) Address() string {
	return fmt.Sprintf("http://%s:%s", s.host, s.port)
}

// Stop _
func (s *Server) Stop() chan struct{} {
	s.stop <- os.Interrupt
	return s.stopped
}

func (s *Server) waitForStop(stop *chan os.Signal, server *http.Server) {
	<-s.stop
	err := s.server.Close()
	if err != nil {
		log.Error.Println(err)
	}
	s.stopped <- struct{}{}
}

// WaitTilRunning _
func (s *Server) WaitTilRunning() {
	<-s.stopped
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "API is running")
}

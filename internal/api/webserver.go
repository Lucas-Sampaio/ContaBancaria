package webserver

import (
	"net/http"

	"github.com/Lucas-Sampaio/ContaBancaria/internal/api/controllers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type WebServer struct {
	router          chi.Router
	WebServerPort   string
	contaController controllers.ContaController
}

func NewWebServer(serverPort string,
	contaController controllers.ContaController) *WebServer {
	return &WebServer{
		router:          chi.NewRouter(),
		WebServerPort:   serverPort,
		contaController: contaController,
	}
}

// register middeleware logger
// start the server
func (s *WebServer) Start() error {
	s.router.Use(middleware.Logger)
	s.router.Use(middleware.Recoverer)
	s.configureRouters()
	err := http.ListenAndServe(s.WebServerPort, s.router)
	return err
}

func (s *WebServer) configureRouters() {
	s.router.Route("/conta", func(r chi.Router) {
		r.Post("/", s.contaController.Create)
	})
}

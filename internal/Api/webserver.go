package webserver

import (
	"net/http"

	"github.com/Lucas-Sampaio/ContaBancaria/internal/Api/controllers"
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
	s.router.Use(JSONMiddleware)

	s.configureRouters()
	err := http.ListenAndServe(s.WebServerPort, s.router)
	return err
}

func (s *WebServer) configureRouters() {
	s.router.Route("/conta", func(r chi.Router) {
		r.Get("/", s.contaController.ObterContas)
		r.Post("/", s.contaController.Criar)
		r.Patch("/{agenciaConta}/desativar", s.contaController.Desativar)
		r.Patch("/{agenciaConta}/bloquear", s.contaController.Bloquear)
		r.Patch("/{agenciaConta}/ativar", s.contaController.Ativar)
	})
}

// JSONMiddleware define o cabeçalho Content-Type para "application/json"
func JSONMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

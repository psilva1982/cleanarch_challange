package webserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type WebServer struct {
	Router        chi.Router
	Handlers      map[string]map[string]http.HandlerFunc // Armazena os handlers por caminho e método
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      make(map[string]map[string]http.HandlerFunc),
		WebServerPort: serverPort,
	}
}

// AddHandler agora aceita o método HTTP, além do caminho e do handler
func (s *WebServer) AddHandler(method, path string, handler http.HandlerFunc) {
	// Inicializa o mapa do caminho se ele ainda não existe
	if s.Handlers[path] == nil {
		s.Handlers[path] = make(map[string]http.HandlerFunc)
	}
	s.Handlers[path][method] = handler
}

// Adiciona middlewares, registra handlers e inicia o servidor
func (s *WebServer) Start() {
	// Middleware de logging
	s.Router.Use(middleware.Logger)

	// Itera sobre os handlers e os registra no roteador com o método e caminho apropriados
	for path, handlersByMethod := range s.Handlers {
		for method, handler := range handlersByMethod {
			switch method {
			case http.MethodGet:
				s.Router.Get(path, handler)
			case http.MethodPost:
				s.Router.Post(path, handler)
			case http.MethodPut:
				s.Router.Put(path, handler)
			case http.MethodDelete:
				s.Router.Delete(path, handler)
			case http.MethodPatch:
				s.Router.Patch(path, handler)
			default:
				s.Router.Method(method, path, handler)
			}
		}
	}

	// Inicia o servidor HTTP
	http.ListenAndServe(s.WebServerPort, s.Router)
}
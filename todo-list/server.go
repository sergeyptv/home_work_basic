package todo

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// объект для работы с сервером
type Server struct {
	httpServer *http.Server
}

// функция для инициализации сервера
func (s *Server) Run(port string, handler *gin.Engine) error {
	s.httpServer = &http.Server{
		Addr:    ":" + port,
		Handler: handler,
	}

	return s.httpServer.ListenAndServe()
}

// функция для завершения работы сервера
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

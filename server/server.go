package server

import (
	"api-temp_conv/server/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type Server struct {
	port string
	server *gin.Engine
}

func NewServer() Server {
	porta := os.Getenv("PORT")

	if porta == "" {
		porta = "5000"
	}
	return Server{
		port:   porta,
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	// Define as rotas
	router := routes.ConfigRoutes(s.server)

	log.Print("servidor rodando na porta: ", s.port)
	log.Fatal(router.Run(":" + s.port))
}
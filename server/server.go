package server

import (
	"log"
	"rest-api/ports"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Server struct {
	handler ports.Handler
}

func NewService(handler ports.Handler) *Server {
	return &Server{handler: handler}
}

// Initialize initializes the server by setting up the routes and handlers for various endpoints.
func (server *Server) Initialize() {
	app := fiber.New()
	app.Use(cors.New())

	log.Fatal(app.Listen(":3000"))
}

package api

import (
	db "github.com/ViniciusReno/fomapi/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/Ingredients", server.addIngredient)
	router.GET("/Ingredients/:id", server.getIngredient)
	router.GET("/Ingredients", server.listIngredient)
	router.GET("/isalive", server.isAlive)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

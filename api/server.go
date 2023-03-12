package api

import (
	db "github.com/deep0ne/bankApp/db/sqlc"
	"github.com/gin-gonic/gin"
)

// HTTP запросы для банковского сервиса
type Server struct {
	store  *db.Store // для взаимодействия с бд
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)

	server.router = router
	return server
}

// старт сервера по адресу
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

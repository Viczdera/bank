package api

import (
	db "github.com/Viczdera/bank/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Server to serve http requests
type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts/all", server.getAccountsList)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errResponse(err error) gin.H {
	return gin.H{"res error": err.Error()}
}

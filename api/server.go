package api

import (
	"fmt"

	db "github.com/Viczdera/bank/db/sqlc"
	"github.com/Viczdera/bank/token"
	"github.com/Viczdera/bank/util"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Server to serve http requests
type Server struct {
	config util.Config
	store  db.Store
	token  token.Maker
	router *gin.Engine
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetrickey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{config: config, store: store, token: tokenMaker}
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	router.POST("/users", server.createUser)
	router.GET("/users/:username", server.getUser)

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	// Use the handler from the other folder
	router.GET("/accounts/all", server.getAccountsList)

	router.POST("/transfer", server.createTransfer)

	server.router = router
	return server, nil
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errResponse(err error) gin.H {
	return gin.H{"res error": err.Error()}
}

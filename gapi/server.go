package gapi

import (
	"fmt"

	db "github.com/Chien179/SimpleBank/db/sqlc"
	"github.com/Chien179/SimpleBank/pb"
	"github.com/Chien179/SimpleBank/token"
	"github.com/Chien179/SimpleBank/util"
	"github.com/gin-gonic/gin"
)

// Server serves gRPC requests for our banking service.
type Server struct {
	pb.UnimplementedSimpleBankServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

// NewServer creates a new gRPC server.
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}

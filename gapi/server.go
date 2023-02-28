package gapi

import (
	"fmt"

	db "github.com/RuhullahReza/simplebank/db/sqlc"
	"github.com/RuhullahReza/simplebank/pb"
	"github.com/RuhullahReza/simplebank/token"
	"github.com/RuhullahReza/simplebank/util"
	"github.com/RuhullahReza/simplebank/worker"
)

type Server struct {
	pb.UnimplementedSimpleBankServer
	config util.Config
	store db.Store
	tokenMaker token.Maker
	taskDistributor worker.TaskDistributor
}

func NewServer(config util.Config, store db.Store, taskDistributor worker.TaskDistributor) (*Server, error){
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config: config,
		store: store,
		tokenMaker: tokenMaker,
		taskDistributor: taskDistributor,
	}


	return server, nil
}
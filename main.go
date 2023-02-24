package main

import (
	"database/sql"
	"log"

	"github.com/RuhullahReza/simplebank/api"
	db "github.com/RuhullahReza/simplebank/db/sqlc"
	"github.com/RuhullahReza/simplebank/util"
	_ "github.com/lib/pq"
)


func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config : ", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to database:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config,store)
	if err != nil {
		log.Fatal("cannot create server: ", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
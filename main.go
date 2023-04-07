package main

import (
	"log"

	"github.com/Pizhlo/go-shop/api"
	"github.com/Pizhlo/go-shop/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// conn, err := sql.Open(config.DBDriver, config.DBSource)
	// if err != nil {
	// 	log.Fatal("cannot connect to DB:", err)
	// }

	//store := db.NewStore(conn)
	server, err := api.NewServer(config)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
	//defer conn.Close()

}

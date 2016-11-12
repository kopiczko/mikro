package main

import (
	"log"

	"github.com/kopiczko/mikro/auth"
	"github.com/kopiczko/mikro/dbaccessor/client"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/server"
)

func main() {
	cmd.Init()
	server.Init(
		server.Name(auth.ServiceName),
	)

	dbAccessor := client.NewDBAccessor()

	server.Handle(server.NewHandler(
		auth.New(dbAccessor),
	))
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"log"

	"github.com/kopiczko/mikro/auth"
	"github.com/kopiczko/mikro/dbaccessor/client"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/server"
)

const ServiceName = "mikro.auth"

func main() {
	cmd.Init()
	server.Init(
		server.Name(ServiceName),
	)
	config := server.DefaultOptions()

	dbAccessor := client.NewDBAccessor(config.Registry)

	server.Handle(server.NewHandler(
		auth.New(dbAccessor),
	))
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

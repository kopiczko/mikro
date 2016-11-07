package main

import (
	"log"

	"github.com/kopiczko/mikro/dbaccessor"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/server"
)

const ServiceName = "mikro.userservice"

func main() {
	cmd.Init()
	server.Init(
		server.Name(ServiceName),
	)
	server.Handle(server.NewHandler(
		new(dbaccessor.DBAccessor),
	))
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

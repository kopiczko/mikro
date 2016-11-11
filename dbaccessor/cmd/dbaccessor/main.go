package main

import (
	"log"

	"github.com/kopiczko/mikro/dbaccessor"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/server"
)

func main() {
	cmd.Init()
	server.Init(
		server.Name(dbaccessor.ServiceName),
	)
	server.Handle(server.NewHandler(
		new(dbaccessor.DBAccessor),
	))
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

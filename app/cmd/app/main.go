package main

import (
	"log"

	"github.com/kopiczko/mikro/app"
	"github.com/kopiczko/mikro/dbaccessor/client"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/server"
)

func main() {
	cmd.Init()
	server.Init(
		server.Name(app.ServiceName),
	)

	dbAccessor := client.NewDBAccessor()

	server.Handle(server.NewHandler(
		app.New(dbAccessor),
	))
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

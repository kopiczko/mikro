package main

import (
	"log"

	"github.com/kopiczko/mikro/dbaccessor"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/server"
)

const UserServiceName = "mikro.userservice"

func main() {
	cmd.Init()
	server.Init(
		server.Name(UserServiceName),
	)
	server.Handle(server.NewHandler(
		new(dbaccessor.UserService),
	))
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

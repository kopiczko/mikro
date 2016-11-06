package main

import (
	"log"

	"github.com/kopiczko/mikro/auth"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/server"
)

const AuthName = "mikro.auth"

func main() {
	cmd.Init()
	server.Init(
		server.Name(AuthName),
	)
	server.Handle(server.NewHandler(
		new(auth.Auth),
	))
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

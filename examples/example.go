package main

import (
	"fmt"
	"log"

	"golang.org/x/net/context"

	appclient "github.com/kopiczko/mikro/app/client"
	authclient "github.com/kopiczko/mikro/auth/client"
	"github.com/micro/go-micro/cmd"
)

func main() {
	cmd.Init()
	app := appclient.NewApp()
	auth := authclient.NewAuth()
	token, err := auth.Login(context.Background(), "pawel", "ignore")
	if err != nil {
		log.Fatalf("Login error: %v", err)
	}
	todoList, err := app.TODOList(context.Background(), token)
	if err != nil {
		log.Fatalf("Login error: %v", err)
	}
	if len(todoList.Items) == 0 {
		fmt.Println("No more to-do items")
	} else {
		fmt.Println("My to-do list:")
	}
	for _, e := range todoList.Items {
		fmt.Println("*", e)
	}
}

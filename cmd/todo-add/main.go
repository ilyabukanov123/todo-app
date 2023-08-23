package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ilyabukanov123/todo-app/internal/app/todo-app/config"
	"github.com/ilyabukanov123/todo-app/internal/app/todo-app/handler"
	"github.com/ilyabukanov123/todo-app/internal/app/todo-app/server"
	"log"
)

const defaultConfigPath = "/Users/ilabukanov/go/src/PetProjects/todo-app/configs/todo-app.json"

func main() {
	confPath := flag.String("config", defaultConfigPath, "")
	flag.Parse()

	config, err := config.GetConfig(*confPath)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(config)

	handlers := new(handler.Handler)

	s := server.NewServer(config, handlers.InitRoutes(), "8080")
	if err := s.Start(); err != nil {
		// throw error
		fmt.Println(err)
	}
	gin.SetMode(gin.ReleaseMode)
}

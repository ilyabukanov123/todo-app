package main

import (
	"flag"
	"fmt"
	"github.com/ilyabukanov123/todo-app/internal/app/todo-app/config"
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

	s := server.NewServer(config)
	if err := s.Start(); err != nil {
		// throw error
		fmt.Println(err)
	}
}

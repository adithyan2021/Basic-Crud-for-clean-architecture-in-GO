package main

import (
	"log"

	"github.com/go-delve/delve/pkg/config"
	"github.com/shakezidin/internal/di"
)

func main() {
	cnf, err := config.LoadConfig()
	if err != nil {
		log.Fatal("failed to load environements")
	}
	server := di.Init(cnf)
	server.StartServer()
}

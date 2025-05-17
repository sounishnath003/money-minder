package main

import (
	"log"

	"github.com/sounishnath003/money-minder/internal/core"
	"github.com/sounishnath003/money-minder/internal/server"
	"github.com/sounishnath003/money-minder/internal/utility"
)

func main() {
	co := &core.Core{
		Port:     utility.GetNumberValueFromEnv("PORT", 3000),
		Hostname: utility.GetStringValueFromEnv("HOSTNAME", "localhost"),
		Logger:   log.Default(),
	}

	srv := server.NewServer(co)
	srv.Start()
}

package main

import (
	"github.com/sounishnath003/money-minder/internal/core"
	"github.com/sounishnath003/money-minder/internal/server"
)

func main() {
	co := core.InitCore()

	srv := server.NewServer(co)
	srv.Start()
}
